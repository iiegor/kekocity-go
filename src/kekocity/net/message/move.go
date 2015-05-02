package message

import (
  "fmt"

  "github.com/bitly/go-simplejson"
)

type MoveMessage struct {
  Despuesfur string
  Eleft int
  Etop int
  Ezindex int
  Harefok string
  Kxa int
  Kya int
  Mabdanum int
  Soobreidp int
  Sorechekenvio string
}

func (mm *MoveMessage) WritePacket() *simplejson.Json {
  // Fake packet
  movePacket := simplejson.New()
  movePacket.Set("idkek", 68880)
  movePacket.Set("trz", mm.Ezindex)
  movePacket.Set("furk", mm.Harefok)
  movePacket.Set("tra", mm.Eleft)
  movePacket.Set("tro", mm.Etop)
  movePacket.Set("my", mm.Kya)
  movePacket.Set("mx", mm.Kxa)
  movePacket.Set("albox", fmt.Sprintf("%vc%v", mm.Kya, mm.Kxa))
  movePacket.Set("mabdanum", mm.Mabdanum)
  movePacket.Set("despuesfur", mm.Despuesfur)
  movePacket.Set("soyrech", mm.Sorechekenvio)

  return MakeMessage(`"hagostopya"`, movePacket)
}
