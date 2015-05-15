package message

import (
  "fmt"

  "github.com/bitly/go-simplejson"
)

type StopMessage struct {
  Idkek int64
  Stand string
  FokStr string
  FokInt int
}

func (sm *StopMessage) WritePacket() *simplejson.Json {
  // Fake packet
  stopPacket := simplejson.New()

  return MakeMessage(fmt.Sprintf(`"golabelnew", "%v", "%v", "%v", %v`, sm.Idkek, sm.Stand, sm.FokStr, sm.FokInt), stopPacket)
}
