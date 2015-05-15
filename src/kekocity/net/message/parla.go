package message

import (
  "fmt"

  "github.com/bitly/go-simplejson"
)

type ParlaMessage struct {
  Idkek int64
  Status string
}

func (pm *ParlaMessage) WritePacket() *simplejson.Json {
  // Fake packet
  parlaPacket := simplejson.New()

  return MakeMessage(fmt.Sprintf(`"parla", "%v", "%v"`, pm.Idkek, pm.Status), parlaPacket)
}
