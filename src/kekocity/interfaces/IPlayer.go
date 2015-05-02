package interfaces

type IPlayer interface {
  GetUsername() string
  GetUserId() int64
  GetCoins() int32

  // Network
  SetNetworkChans(input <-chan interface{})
}
