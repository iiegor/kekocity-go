package message

import (
  "github.com/bitly/go-simplejson"
)

type AuthMessage struct {
  Status string
  Estado string
  Tipocuenta int
  Plac string
  Tamano int
  Cuantosonnn int
  Imagenmini string
  Id int64
  Keko string
  Creditos int32
  Fichas int32
  Vip int
  Aspectonew string
}

func (am *AuthMessage) WritePacket() *simplejson.Json {
  // Fake packet
  authPacket := simplejson.New()
  authPacket.Set("estado", 1)
  authPacket.Set("tipocuenta", 0)
  authPacket.Set("plac", "mod4")
  authPacket.Set("cuantosonnn", 101)
  authPacket.Set("imagenmini", "02/f/a/3/c/cfa3cae4986d04c03a75e4dbac3db23d.jpg")
  authPacket.Set("id", am.Id)
  authPacket.Set("keko", am.Keko)
  authPacket.Set("creditos", am.Creditos)
  authPacket.Set("status", am.Status)
  authPacket.Set("fichas", am.Fichas)
  authPacket.Set("vip", 34)
  authPacket.Set("asceptonew", "LLL31LLL17/34L5/34LLL17/34L4/1LL2")
  authPacket.Set("tamano", 0)

  return MakeMessage(`"ready"`, authPacket)
}
