package message

import (
  "github.com/bitly/go-simplejson"
)

type AuthMessage struct {
  Status string
}

func (am *AuthMessage) WritePacket() *simplejson.Json {
  authPacket := simplejson.New()
  authPacket.Set("status", am.Status)

  return MakeMessage("auth", authPacket)
}
