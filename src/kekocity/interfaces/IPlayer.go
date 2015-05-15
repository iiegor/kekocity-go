package interfaces

type IPlayer interface {
  GetUsername() string
  GetPlayerId() int64
  GetCoins() int32
  GetClouds() int32
}
