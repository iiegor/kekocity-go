/*
 * TODO: Parse the incoming packet, receivepoller, sendpoller, newconnection
 */
package net

import (
  "log"

  "golang.org/x/net/websocket"

  pnet "kekocity/misc/packet"
  "kekocity/interfaces"
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
  go connection.SendPoller()

  return connection
}

func (c *Connection) AssignUser(_user interfaces.IUser) {
  if _user == nil {
    panic("net.connection: the user interface can not be nil!")
    return
  }

  c.user = _user
  _user.SetNetworkChans(c.rxChan, c.txChan)
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
  for {
    // Read messages from transmit channel
    message := <-c.txChan

    if message == nil {
      log.Println("SenPoller", "The message is nil, break the loop")
      break
    }

    // Convert netmessage to packet
    packet := message.WritePacket()
    packet.SetHeader()

    // Create byte buffer
    buffer := packet.GetBuffer()
    data := buffer[0:packet.GetMsgSize()]

    // Send bytes off to the internetz
    websocket.Message.Send(c.socket, data)
  }
}

func (c *Connection) parsePacket(_packet pnet.IPacket) {
  log.Println("net.connection:", "Received new packet!")
}

func (c *Connection) Close() {
  // Close channels
  close(c.txChan)
  close(c.rxChan)

  // Close the socket
  c.socket.Close()

  c.user = nil
}
