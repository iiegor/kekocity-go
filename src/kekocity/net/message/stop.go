package message

import (
  "fmt"

  "github.com/bitly/go-simplejson"
)

type StopMessage struct {
  Stand string
  FokStr string
  FokInt int
}

func (sm *StopMessage) WritePacket() *simplejson.Json {
  // Fake packet
  stopPacket := simplejson.New()

  return MakeMessage(fmt.Sprintf(`"golabelnew", 68880, "%v", "%v", %v`, sm.Stand, sm.FokStr, sm.FokInt), stopPacket)
}
