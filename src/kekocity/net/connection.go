package net

import (
  "fmt"
  "errors"

  "github.com/gorilla/websocket"
  "github.com/bitly/go-simplejson"

  "kekocity/interfaces"
  netmsg "kekocity/net/message"
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

  player interfaces.IPlayer
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

func (c *Connection) AssignToPlayer(_player interfaces.IPlayer) {
  if _player == nil {
    panic("Connection - Player interface can not be nil")
  }

  c.player = _player
  _player.SetNetworkChans(c.output)
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
  namespace, err := obj.GetIndex(0).String()
  if len(namespace) < 1 || err != nil {
    fmt.Println("Namespace can not be nil")
    return
  }

  switch namespace {
  case "auth":
    player, err := netmsg.AuthPacket(obj.GetIndex(1))
    if err != nil {
      c.Close()
      return
    }

    c.AssignToPlayer(player)
  default:
    fmt.Printf("Unhandled packet received - %v\n", namespace)
  }
}

func (c *Connection) Close() {
  // Close the websocket
  c.ws.Close()

  c.player = nil
}
