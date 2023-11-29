package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
	"log"
	"net"
	"time"
)

const (
	ttl     = time.Second * 8
	checkId = "check_health"
)

type ConsulService struct {
	consulClient *api.Client
}

func NewConsulService() *ConsulService {
	client, err := api.NewClient(&api.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &ConsulService{
		consulClient: client,
	}
}

func (s *ConsulService) start() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	s.registerConsul()
	go s.updateHealthCheck()
	s.acceptLoop(ln)
}

func (s *ConsulService) acceptLoop(ln net.Listener) {
	for {
		_, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s *ConsulService) updateHealthCheck() {
	ticker := time.NewTicker(time.Second * 5)
	for {
		err := s.consulClient.Agent().UpdateTTL(checkId, "online", api.HealthPassing)
		if err != nil {
			log.Fatal(err)
		}

		<-ticker.C
	}
}

func (s *ConsulService) registerConsul() {
	check := &api.AgentServiceCheck{
		DeregisterCriticalServiceAfter: ttl.String(),
		TLSSkipVerify:                  true,
		TTL:                            ttl.String(),
		CheckID:                        checkId,
	}

	register := &api.AgentServiceRegistration{
		ID:      "login_service",
		Name:    "mycluster",
		Tags:    []string{"login"},
		Address: "127.0.0.1",
		Port:    3000,
		Check:   check,
	}

	query := map[string]any{
		"type":        "service",
		"service":     "mycluster",
		"passingonly": true,
	}

	plan, err := watch.Parse(query)
	if err != nil {
		log.Fatal(err)
	}

	plan.HybridHandler = func(index watch.BlockingParamVal, result any) {
		switch msg := result.(type) {
		case []*api.ServiceEntry:
			for _, entry := range msg {
				fmt.Println("new member joined", entry.Service)
			}
		}
	}
	go func() {
		err := plan.RunWithConfig("", &api.Config{})
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err := s.consulClient.Agent().ServiceRegister(register); err != nil {
		log.Fatal(err)
	}
}

func main() {
	s := NewConsulService()
	s.start()
}
