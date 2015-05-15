package models

import (
  "kekocity/data/entities"
  "kekocity/interfaces"
)

type Room struct {
  RoomEntity *entities.Room

  Players map[interfaces.IPlayer]bool
}

func NewRoom(_entity *entities.Room) *Room {
  r := &Room{}
  r.RoomEntity = _entity
  r.Players = make(map[interfaces.IPlayer]bool)

  return r
}

// Room methods
func (r *Room) GetId() int {
	return r.RoomEntity.Id
}

func (r *Room) GetName() string {
	return r.RoomEntity.Nombre
}

func (r *Room) GetPlayers() map[interfaces.IPlayer]bool {
  return r.Players
}

// !!
func (r *Room) GetEntity() *entities.Room {
  return r.RoomEntity
}
