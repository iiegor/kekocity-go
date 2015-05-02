package message

import (
  "github.com/bitly/go-simplejson"
)

type DonveMessage struct {
}

func (am *DonveMessage) WritePacket() *simplejson.Json {
  // Fake packet
  donvePacket := simplejson.New()
  donvePacket.Set("conpowas", 2)
  donvePacket.Set("bancasa", 0)
  donvePacket.Set("id", 152964)
  donvePacket.Set("pass", "pp")
  donvePacket.Set("idkeko", 68880)
  donvePacket.Set("maxgente", 25)
  donvePacket.Set("kekosonline", 0)
  donvePacket.Set("pais", "es")
  donvePacket.Set("empresa", 0)
  donvePacket.Set("nombre", "Mi sala")
  donvePacket.Set("descripcion", "")
  donvePacket.Set("permisos", 0)
  donvePacket.Set("cuadrosy", 10)
  donvePacket.Set("cuadrosx", 10)
  donvePacket.Set("keko", "undefined")
  donvePacket.Set("tipo", 1)
  donvePacket.Set("estilopared", 9)
  donvePacket.Set("estiloparedd", 9)
  donvePacket.Set("estilosuelo", 135)
  donvePacket.Set("diagonal", 0)
  donvePacket.Set("cambios", 0)
  donvePacket.Set("notraspasar", 0)
  donvePacket.Set("armas", 1)
  donvePacket.Set("aspectonew", "LLL4LLLLLLLLLL2")
  donvePacket.Set("dueno", 0)
  donvePacket.Set("mirployy", -9)
  donvePacket.Set("mirploxx", -10)
  donvePacket.Set("sy", 0)
  donvePacket.Set("sx", 0)
  donvePacket.Set("solopongos", "no")

  return MakeMessage(`"ventryacarg", 152964`, donvePacket)
}
