package message

import (
  "github.com/bitly/go-simplejson"
)

type ClientMessage struct {
}

func (am *ClientMessage) WritePacket() *simplejson.Json {
  // Fake packet
  clientPacket := simplejson.New()
  clientPacket.Set("idkek", 68880)
  clientPacket.Set("ecam", "")
  clientPacket.Set("kek", "undefined")
  clientPacket.Set("py", 0)
  clientPacket.Set("px", 2)
  clientPacket.Set("cabez", "")
  clientPacket.Set("conpowas", 2)
  clientPacket.Set("ceco", "3")
  clientPacket.Set("csoc", "3")
  clientPacket.Set("megustas", 17)
  clientPacket.Set("fok", "standL4LLL")
  clientPacket.Set("parlando", "")
  clientPacket.Set("tags", "")
  clientPacket.Set("vida", 100)
  clientPacket.Set("rop", "LLL4LLLLLLLLLL2")
  clientPacket.Set("tam", 0)
  clientPacket.Set("placs", "")
  clientPacket.Set("imagenmini", "02/4/0/0/c/a400c38a79167027f543d3542faeb74d.jpg")
  clientPacket.Set("sobreidp", 0)
  clientPacket.Set("misi", "")
  clientPacket.Set("tipocuenta", 0)

  return MakeMessage(`"roomclientsya"`, clientPacket)
}
