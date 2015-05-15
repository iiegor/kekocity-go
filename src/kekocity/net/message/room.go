package message

import (
  "fmt"

  "github.com/bitly/go-simplejson"
)

type RoomMessage struct {
  Id string
  Conpowas int
  Bancasa int
  Idkeko int64
  Maxgente int
  Kekosonline int
  Nombre string
  Descripcion string
  Permisos int
  Cuadrosy int
  Cuadrosx int
  Keko string
  Tipo int
  Estilopared int
  Estiloparedd int
  Estilosuelo int
  Diagonal int
  Cambios int
  Notraspasar int
  Armas int
  Aspectonew string
  Dueno int
  Mirployy int
  Mirploxx int
  Sy int
  Sx int
  Solopongos string
  Pais int
  Empresa int
}

func (rm *RoomMessage) WritePacket() *simplejson.Json {
  roomPacket := simplejson.New()
  roomPacket.Set("conpowas", rm.Conpowas)
  roomPacket.Set("bancasa", rm.Bancasa)
  roomPacket.Set("id", rm.Id)
  roomPacket.Set("pass", "pp")
  roomPacket.Set("idkeko", rm.Idkeko)
  roomPacket.Set("maxgente", rm.Maxgente)
  roomPacket.Set("kekosonline", rm.Kekosonline)
  roomPacket.Set("pais", rm.Pais)
  roomPacket.Set("empresa", rm.Empresa)
  roomPacket.Set("nombre", rm.Nombre)
  roomPacket.Set("descripcion", rm.Descripcion)
  roomPacket.Set("permisos", rm.Permisos)
  roomPacket.Set("cuadrosy", rm.Cuadrosy)
  roomPacket.Set("cuadrosx", rm.Cuadrosx)
  roomPacket.Set("keko", rm.Keko)
  roomPacket.Set("tipo", rm.Tipo)
  roomPacket.Set("estilopared", rm.Estilopared)
  roomPacket.Set("estiloparedd", rm.Estiloparedd)
  roomPacket.Set("estilosuelo", rm.Estilosuelo)
  roomPacket.Set("diagonal", rm.Diagonal)
  roomPacket.Set("cambios", rm.Cambios)
  roomPacket.Set("notraspasar", rm.Notraspasar)
  roomPacket.Set("armas", rm.Armas)
  roomPacket.Set("aspectonew", rm.Aspectonew)
  roomPacket.Set("dueno", rm.Dueno)
  roomPacket.Set("mirployy", rm.Mirployy)
  roomPacket.Set("mirploxx", rm.Mirploxx)
  roomPacket.Set("sy", rm.Sy)
  roomPacket.Set("sx", rm.Sx)
  roomPacket.Set("solopongos", rm.Solopongos)

  return MakeMessage(fmt.Sprintf(`"ventryacarg", %v`, rm.Id), roomPacket)
}
