package model

import (
	"dartsvader-go/helper"

	"github.com/google/uuid"
)

const (
	StatusCreate  int = 0
	StatusStarted int = 1
)

type Game struct {
	ID            uuid.UUID       `json:"id"`
	Name          string          `json:"name"`
	Players       map[int]*Player `json:"players"`
	Status        int             `json:"status"`
	currentPlayer int
	maxPlayer     int
}

func NewGame() *Game {
	return &Game{
		ID:            uuid.New(),
		Status:        StatusCreate,
		currentPlayer: 0,
		maxPlayer:     2,
		Name:          "501",
		Players:       make(map[int]*Player, 2),
	}
}

func (g *Game) SetPlayer(player *Player) {
	pos := len(g.Players)
	g.Players[pos] = player
}

func (g *Game) GetCurrentPlayer() *Player {
	return g.Players[g.currentPlayer]
}

func (g *Game) GetCurrentPlayerId() int {
	return g.currentPlayer
}

func (g *Game) NextPlayer() {
	_, _, next := helper.GetMapPosition(g.Players, g.currentPlayer)
	g.currentPlayer = next
}