package net

import (
  "fmt"
  "errors"

  "github.com/gorilla/websocket"
  "github.com/bitly/go-simplejson"

  "kekocity/interfaces"
  //netmsg "kekocity/net/message"
)

const (
	fps  = 24               // Frames per second.
	fpsl = 1000 / fps       // Duration of a single (milliseconds)
	fpsn = 1000000000 / fps // Duration of a single frame (nanoseconds)
)

var ErrNoWebsocket = errors.New(`Don't have any websocket connection.`)

type Connection struct {
  ws *websocket.Conn

  output chan *simplejson.Json

  user interfaces.IUser
}

func NewConnection(_ws *websocket.Conn) *Connection {
  // The pointer allow us to modify connection struct from outside
  connection := &Connection{
    ws: _ws,
    output: make(chan *simplejson.Json),
  }

  go connection.Writer()
  connection.Reader()

  return connection
}

func (c *Connection) Writer() {
  for {
    // Read messages from transmit channel
    netmessage := <-c.output

    if netmessage == nil {
      fmt.Println("ConnectionWrapper", "SendPoller", "Netmessage == nil, breaking loop")
      break
    }

    // Send bytes off to the internetz
    c.ws.WriteJSON(netmessage)
	}
}

func (c *Connection) Reader() {
  for {
    _, message, err := c.ws.ReadMessage()
    if err != nil {
      break
    }

    obj, err := simplejson.NewJson([]byte(message))
    c.processPacket(obj)
	}
}

func (c *Connection) processPacket(obj *simplejson.Json) {
  c.output <- obj
}

func (c *Connection) Close() {
  // Close the websocket
  c.ws.Close()

  c.user = nil
}
