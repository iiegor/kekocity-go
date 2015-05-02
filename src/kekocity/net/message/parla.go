package message

import (
  "fmt"

  "github.com/bitly/go-simplejson"
)

type ParlaMessage struct {
  Status string
}

func (pm *ParlaMessage) WritePacket() *simplejson.Json {
  // Fake packet
  parlaPacket := simplejson.New()

  return MakeMessage(fmt.Sprintf(`"parla", 68880, "%v"`, pm.Status), parlaPacket)
}
