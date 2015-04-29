package models

import (
  "fmt"

  "github.com/bitly/go-simplejson"

  "kekocity/data/entities"
)

type Player struct {
	PlayerEntity *entities.Player

	output <-chan *simplejson.Json
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
func (p *Player) SetNetworkChans(_output <-chan *simplejson.Json) {
	p.output = _output

	go p.netReceiveMessages()
}

func (p *Player) netReceiveMessages() {
  for {
    message := <-p.output
    if message == nil {
      break
    }

    fmt.Println("Received from player's netReceiveMessages:", message)
  }
}
