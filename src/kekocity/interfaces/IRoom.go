package interfaces

import "kekocity/data/entities"

type IRoom interface {
  GetName() string
  GetId() int
  GetPlayers() map[IPlayer]bool

  GetEntity() *entities.Room
}
