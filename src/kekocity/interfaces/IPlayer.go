package interfaces

import "github.com/bitly/go-simplejson"

type IPlayer interface {
  GetUsername() string
  GetUserId() int64
  GetCoins() int32

  // Network
  SetNetworkChans(output <-chan *simplejson.Json)
}
