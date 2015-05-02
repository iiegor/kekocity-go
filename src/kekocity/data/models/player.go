package models

import (
  "fmt"

  "kekocity/data/entities"
)

type Player struct {
	PlayerEntity *entities.Player

	input <-chan interface{}
}

func NewPlayer(_entity *entities.Player) *Player {
  p := &Player{}
  p.PlayerEntity = _entity

  return p
}

// User methods
func (p *Player) GetUsername() string {
	return p.PlayerEntity.Username
}

func (p *Player) GetUserId() int64 {
	return int64(p.PlayerEntity.Id)
}

func (p *Player) GetCoins() int32 {
	return p.PlayerEntity.Coins
}

// Network
func (p *Player) SetNetworkChans(_input <-chan interface{}) {
	p.input = _input

	go p.netReceiveMessages()
}

func (p *Player) netReceiveMessages() {
  for {
    message := <-p.input

    fmt.Println("Received from player's netReceiveMessages:", message)
  }
}
