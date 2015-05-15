package models

import (
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

func (p *Player) GetPlayerId() int64 {
	return int64(p.PlayerEntity.Id)
}

func (p *Player) GetCoins() int32 {
	return p.PlayerEntity.Coins
}

func (p *Player) GetClouds() int32 {
  return p.PlayerEntity.Clouds
}
