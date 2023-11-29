package main

import (
	"fmt"
)

type Player struct {
	Name string
}

type GameState struct {
	players []*Player

	msgch chan any
}

func (g *GameState) Receive(msg any) {
	g.msgch <- msg
}

func (g *GameState) loop() {
	for msg := range g.msgch {
		g.handleMessage(msg)
	}
}

func (g *GameState) handleMessage(msg any) {
	switch v := msg.(type) {
	case *Player:
		g.addPlayer(v)
	default:
		panic("invalid message received")
	}
}

func (g *GameState) addPlayer(p *Player) {
	g.players = append(g.players, p)

	fmt.Println("adding player:", p.Name)
}

func NewGameState() *GameState {
	g := &GameState{
		players: []*Player{},
		msgch:   make(chan any, 10),
	}
	go g.loop()

	return g
}

type MutexServer struct {
	gameState *GameState
}

func NewMutexServer() *MutexServer {
	return &MutexServer{
		gameState: NewGameState(),
	}
}

func (s *MutexServer) handleNewPlayer(player *Player) error {
	s.gameState.Receive(player)
	return nil
}
