package net

import (
  "log"

  "golang.org/x/net/websocket"

  pnet "kekocity/misc/packet"
  "kekocity/interfaces"
  //"kekocity/net/message"
)

type Connection struct {
  socket *websocket.Conn

  txChan chan pnet.INetMessageWriter
	rxChan chan pnet.INetMessageReader

  user interfaces.IUser
}

func NewConnection(_socket *websocket.Conn) *Connection {
  // The pointer allow us to modify connection struct from outside
  connection := &Connection{
    socket: _socket,
    txChan: make(chan pnet.INetMessageWriter),
    rxChan: make(chan pnet.INetMessageReader),
  }

  go connection.ReceivePoller()

  return connection
}

func (c *Connection) AssignToUser(_user interfaces.IUser) {
  if _user == nil {
    panic("net.connection: the user interface can not be nil!")
    return
  }

  c.user = _user
  _user.SetNetworkChans(c.rxChan, c.txChan)
}

func (c *Connection) ReceivePoller() {
  for {
    packet := pnet.NewPacket()
		var buffer []uint8
    err := websocket.Message.Receive(c.socket, &buffer)

    if err == nil {
			copy(packet.Buffer[0:len(buffer)], buffer[0:len(buffer)])

			c.processPacket(packet)
		} else {
			println(err.Error())
			break
		}
  }
}

func (c *Connection) processPacket(_packet pnet.IPacket) {
  log.Println("Received packet:", _packet.ToString())

  // Test response
  websocket.JSON.Send(c.socket, c.user);
}

func (c *Connection) Close() {
  // Close channels
  close(c.txChan)
  close(c.rxChan)

  // Close the socket
  c.socket.Close()

  c.user = nil
}
