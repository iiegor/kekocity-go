package message

import pnet "kekocity/misc/packet"

type AuthMessage struct {
  Status string
}

func (m *AuthMessage) WritePacket() pnet.IPacket {
  packet := pnet.NewPacket()
  packet.AddString(string(m.Status))

  return packet
}
