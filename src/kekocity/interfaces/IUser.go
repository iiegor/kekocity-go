package interfaces

import pnet "kekocity/misc/packet"

type IUser interface {
  GetUsername() string
  GetUserId() int64
  GetCoins() int32

  // Network
  SetNetworkChans(_rx <-chan pnet.INetMessageReader, _tx chan<- pnet.INetMessageWriter)
}
