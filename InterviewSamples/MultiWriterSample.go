package main

import (
	"bytes"
	"fmt"
	"io"
)

type MWConn struct {
	io.Writer
}

func NewMWConn() *MWConn {
	return &MWConn{
		Writer: new(bytes.Buffer),
	}
}

func (c *MWConn) Write(b []byte) (int, error) {
	fmt.Println("writing to underlying connection:", string(b))

	return c.Writer.Write(b)
}

type MWServer struct {
	peers map[*MWConn]bool
}

func NewMWServer() *MWServer {
	s := &MWServer{
		peers: make(map[*MWConn]bool),
	}

	for i := 0; i < 10; i++ {
		s.peers[NewMWConn()] = true
	}

	return s
}

func (s *MWServer) broadcast(msg []byte) error {
	peers := []io.Writer{}

	for peer := range s.peers {
		peers = append(peers, peer)
	}

	mw := io.MultiWriter(peers...)
	_, err := mw.Write(msg)
	return err
}

func main() {
	s := NewMWServer()
	s.broadcast([]byte("foo"))
}
