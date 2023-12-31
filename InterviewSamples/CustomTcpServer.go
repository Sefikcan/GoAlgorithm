package main

import (
	"fmt"
	"log"
	"net"
)

type server struct {
	listenAddress string
	ln            net.Listener
	quitch        chan struct{}
	msgch         chan []byte
}

func newServer(listenAddress string) *server {
	return &server{
		listenAddress: listenAddress,
		quitch:        make(chan struct{}),
		msgch:         make(chan []byte, 10),
	}
}

func (s *server) start() error {
	ln, err := net.Listen("tcp", s.listenAddress)
	if err != nil {
		return err
	}
	defer ln.Close()

	s.ln = ln

	go s.acceptLoop()

	<-s.quitch
	close(s.msgch)

	return nil
}

func (s *server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		fmt.Println("new connection to the server:", conn.RemoteAddr())

		go s.readLoop(conn)
	}
}

func (s *server) readLoop(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}

		s.msgch <- buf[:n]
	}
}

func main() {
	server := newServer(":3000")

	go func() {
		for msg := range server.msgch {
			fmt.Println("received message from connection:", string(msg))
		}
	}()

	log.Fatal(server.start())
}
