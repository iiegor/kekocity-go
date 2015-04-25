package interfaces

import pnet "kekocity/misc/packet"

type IPlayer interface {
  GetPlayerId() int64

  // Network
  SetNetworkChans(_rx <-chan pnet.INetMessageReader, _tx chan<- pnet.INetMessageWriter)
}
