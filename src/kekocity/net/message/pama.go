package message

import (
  "fmt"

  "github.com/bitly/go-simplejson"
)

type PamaMessage struct {
  Idkek int64
  Color int
  Message string
  Type int
}

func (pm *PamaMessage) WritePacket() *simplejson.Json {
  // Fake packet
  pamaPacket := simplejson.New()

  return MakeMessage(fmt.Sprintf(`"pamus", %v, "", %v, "%v", %v`, pm.Idkek, pm.Color, pm.Message, pm.Type), pamaPacket)
}
