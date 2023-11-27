package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
	"time"
)

type SocketServer struct {
	conns map[*websocket.Conn]bool
}

func NewSocketServer() *SocketServer {
	return &SocketServer{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *SocketServer) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client:", ws.RemoteAddr())

	s.conns[ws] = true

	s.readLoop(ws)
}

func (s *SocketServer) handleWSRecipes(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client to recipes feed:", ws.RemoteAddr())

	for {
		payload := fmt.Sprintf("recipe data -> %d\n", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(time.Second * 2)
	}
}

func (s *SocketServer) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			continue
		}

		msg := buf[:n]
		s.broadcast(msg)
	}
}

func (s *SocketServer) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("write error", err)
			}
		}(ws)
	}
}

func main() {
	server := NewSocketServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/ws/recipes", websocket.Handler(server.handleWSRecipes))
	http.ListenAndServe(":1341", nil)
}

/*
Browser test js code
let socket = new WebSocket("ws://localhost:1341/ws");
socket.onmessage = (event) => { console.log("received from the server", event.data) }
socket.send("hello from client")
*/
