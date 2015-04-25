package models

import (
  "log"

  "github.com/eaigner/hood"

  pnet "kekocity/misc/packet"
  "kekocity/data/entities"
)

type User struct {
  Username string
	UserEntity *entities.User

	db *hood.Hood

	rxChan <-chan pnet.INetMessageReader
	txChan chan<- pnet.INetMessageWriter
}

func NewUser(_entity *entities.User, _db *hood.Hood) *User {
  u := &User{}
  u.UserEntity = _entity
  u.db = _db

  return u
}

// User methods
func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetUserId() int64 {
	return int64(u.UserEntity.Id)
}

func (u *User) GetCoins() int32 {
	return u.UserEntity.Coins
}

// Network
func (u *User) SetNetworkChans(_rx <-chan pnet.INetMessageReader, _tx chan<- pnet.INetMessageWriter) {
	u.rxChan = _rx
	u.txChan = _tx

	go u.netReceiveMessages()
}

func (u *User) netReceiveMessages() {
  for {
    message := <-u.rxChan
    if message == nil {
      break
    }

    log.Println("Received:", message)
  }
}
