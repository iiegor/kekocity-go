/*
 * TODO: Parse the incoming packet, receivepoller, sendpoller, newconnection
 */
package net

import (
  "log"

  "golang.org/x/net/websocket"

  pnet "kekocity/misc/packet"
)

type Connection struct {
  socket *websocket.Conn

  txChan chan pnet.INetMessageWriter
	rxChan chan pnet.INetMessageReader
}

func NewConnection(_socket *websocket.Conn) *Connection {
  // The pointer allow us to modify connection struct from outside
  connection := &Connection{
    socket: _socket,
    txChan: make(chan pnet.INetMessageWriter),
    rxChan: make(chan pnet.INetMessageReader),
  }

  go connection.ReceivePoller()
  go connection.SendPoller()

  return connection
}

/*
 * ReceivePoller and SendPoller starts listening when the first packet is verified and the new connection is started
 */
func (c *Connection) ReceivePoller() {
  for {
    packet := pnet.NewPacket()

    var buffer []uint8
		err := websocket.Message.Receive(c.socket, &buffer)

		if err == nil {
			copy(packet.Buffer[0:len(buffer)], buffer[0:len(buffer)])

			c.parsePacket(packet)
		} else {
			println(err.Error())
			break
		}
  }
}

func (c *Connection) SendPoller() {
}

func (c *Connection) parsePacket(_packet pnet.IPacket) {
  log.Println(_packet)
}

func (c *Connection) Close() {
  // Close channels
  close(c.txChan)
  close(c.rxChan)

  // Close the socket
  c.socket.Close()
}
