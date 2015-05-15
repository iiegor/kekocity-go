package message

import (
  "fmt"

  "github.com/bitly/go-simplejson"
)

type CmdMessage struct {
  User string
  Cmd string
  Param string
}

func (cm *CmdMessage) WritePacket() *simplejson.Json {
  // Fake packet
  cmdPacket := simplejson.New()

  return MakeMessage(fmt.Sprintf(`"avisos", "%v", "%v", "%v"`, cm.User, cm.Cmd, cm.Param), cmdPacket)
}
