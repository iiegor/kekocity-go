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

  // Network channels
  output chan interface{}
  input chan interface{}

  player interfaces.IPlayer
}

func NewConnection(_ws *websocket.Conn) *Connection {
  // The pointer allow us to modify connection struct from outside
  connection := &Connection{
    ws: _ws,
    output: make(chan interface{}),
    input: make(chan interface{}),
  }

  go connection.Writer()
  go connection.Reader()

  return connection
}

func (c *Connection) AssignToPlayer(_player interfaces.IPlayer) {
  if _player == nil {
    panic("Connection - Player interface can not be nil")
  }

  c.player = _player
  /* TODO: Issue #2 */
  _player.SetNetworkChans(c.input)
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
    if err != nil {
      fmt.Println("Can not parse the message:", string(message))
      continue
    }

    c.processPacket(obj)
	}
}

func (c *Connection) processPacket(obj *simplejson.Json) {
  namespace, err := obj.GetIndex(0).String()
  if len(namespace) < 1 || err != nil {
    fmt.Println("Namespace can not be nil:", obj)
    return
  }

  /* TODO: See issue #2 */
  // Unauth packet handler
  switch namespace {
    case "donve":
      donveMessage := &netmsg.DonveMessage{}

      c.output <- donveMessage.WritePacket()
    case "ventryatoy":
      clientMessage := &netmsg.ClientMessage{}

      c.output <- clientMessage.WritePacket()
    case "parlo":
      parloStatus, _ := obj.GetIndex(1).String()
      parlaMessage := &netmsg.ParlaMessage{}

      if parloStatus == "habla" {
        parlaMessage.Status = "habla"
      } else {
        parlaMessage.Status = "n"
      }

      c.output <- parlaMessage.WritePacket()
    case "pama":
      c.input <- obj
    case "hagostopnew":
      moveInfo := obj.GetIndex(1)
      despuesfur, _ := moveInfo.Get("despuesfur").String()
      eleft, _ := moveInfo.Get("eleft").Int()
      etop, _ := moveInfo.Get("etop").Int()
      ezindex, _ := moveInfo.Get("ezindex").Int()
      harefok, _ := moveInfo.Get("harefok").String()
      kxa, _ := moveInfo.Get("kxa").Int()
      kya, _ := moveInfo.Get("kya").Int()
      mabdanum, _ := moveInfo.Get("mabdanum").Int()
      soobreidp, _ := moveInfo.Get("soobreidp").Int()
      sorechekenvio, _ := moveInfo.Get("sorechekenvio").String()

      moveMessage := &netmsg.MoveMessage{
        Despuesfur: despuesfur,
        Eleft: eleft,
        Etop: etop,
        Ezindex: ezindex,
        Harefok: harefok,
        Kxa: kxa,
        Kya: kya,
        Mabdanum: mabdanum,
        Soobreidp: soobreidp,
        Sorechekenvio: sorechekenvio,
      }

      c.output <- moveMessage.WritePacket()
    case "golabelnew":
      stand, _ := obj.GetIndex(1).String()
      fokStr, _ := obj.GetIndex(2).String()
      fokInt, _ := obj.GetIndex(3).Int()

      stopMessage := &netmsg.StopMessage{
        Stand: stand,
        FokStr: fokStr,
        FokInt: fokInt,
      }

      c.output <- stopMessage.WritePacket()
    default:
      fmt.Printf("Unhandled packet received - %v\n", namespace)
  }
}

func (c *Connection) Close() {
  // Close the websocket
  c.ws.Close()

  c.player = nil
}
