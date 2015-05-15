package net

import (
  "fmt"
  "errors"
  "strconv"

  "github.com/gorilla/websocket"
  "github.com/bitly/go-simplejson"

  "kekocity/interfaces"
  "kekocity/data/helpers"

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
  broadcast chan interface{}
  output chan interface{}
  input chan interface{}

  player interfaces.IPlayer
}
type hub struct {
  connections map[*Connection]bool
  rooms map[string]*netmsg.RoomMessage
}

func newHub() *hub {
  return &hub{
    connections: make(map[*Connection]bool),
    rooms: make(map[string]*netmsg.RoomMessage),
  }
}

var Hub = newHub()

func NewConnection(_ws *websocket.Conn) *Connection {
  // The pointer allow us to modify connection struct from outside
  connection := &Connection{
    ws: _ws,
    broadcast: make(chan interface{}),
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

  Hub.connections[c] = true

  c.player = _player
}

func (c *Connection) Writer() {
  for {
    select {
      case netmessage := <-c.output:
        if netmessage == nil {
          fmt.Println("Connection", "Writer", "Netmessage == nil, breaking loop")
          break
        }
        c.ws.WriteJSON(netmessage)
      case netmessage := <-c.broadcast:
        if netmessage == nil {
          fmt.Println("Connection", "Writer", "Netmessage == nil, breaking loop")
          break
        }

        for hc := range Hub.connections {
          hc.ws.WriteJSON(netmessage)
        }
    }
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

  switch namespace {
    case "donve":
      // Write room packet
      var roomMessage *netmsg.RoomMessage
      roomId, _ := obj.GetIndex(1).String()

      // Check if room is already cached
      if room := Hub.rooms[roomId]; room != nil {
        roomMessage = room
      } else {
        room, err := helpers.RoomHelper.EnterRoom(roomId)
        if err != nil {
          fmt.Println("Room not found, need to send a message back!")
          return
        }

        roomData := room.GetEntity()
        roomMessage = &netmsg.RoomMessage{
          Id: roomId,
          Conpowas: 0,
          Bancasa: 0,
          Idkeko: c.player.GetPlayerId(),
          Maxgente: roomData.Maxgente,
          Kekosonline: 0,
          Nombre: roomData.Nombre,
          Descripcion: roomData.Descripcion,
          Permisos: 0,
          Cuadrosy: roomData.Cuadrosy,
          Cuadrosx: roomData.Cuadrosx,
          Keko: c.player.GetUsername(),
          Tipo: roomData.Tipo,
          Estilopared: roomData.Estilopared,
          Estiloparedd: roomData.Estiloparedd,
          Estilosuelo: roomData.Estilosuelo,
          Diagonal: roomData.Diagonal,
          Cambios: roomData.Cambios,
          Notraspasar: roomData.Notraspasar,
          Armas: roomData.Armas,
          Aspectonew: roomData.Aspectonew,
          Dueno: roomData.Dueno,
          Mirployy: roomData.Mirployy,
          Mirploxx: roomData.Mirploxx,
          Sy: roomData.Sy,
          Sx: roomData.Sx,
          Solopongos: roomData.Solopongos,
          Pais: roomData.Pais,
          Empresa: roomData.Empresa,
        }

        // Push to cache
        Hub.rooms[roomId] = roomMessage
      }

      c.output <- roomMessage.WritePacket()
    case "ventryatoy":
      // Send items in the room
      // Send clients in the room
      clientMessage := &netmsg.ClientMessage{
        Idkek: c.player.GetPlayerId(),
        Kek: c.player.GetUsername(),
      }

      // Send user entity
      c.output <- clientMessage.WritePacket()
    case "parlo":
      parloStatus, _ := obj.GetIndex(1).String()
      parlaMessage := &netmsg.ParlaMessage{
        Idkek: c.player.GetPlayerId(),
      }

      if parloStatus == "habla" {
        parlaMessage.Status = "habla"
      } else {
        parlaMessage.Status = "n"
      }

      c.broadcast <- parlaMessage.WritePacket()
    case "pama":
      pamaColor, _ := obj.GetIndex(1).Int()
      pamaMsg, _ := obj.GetIndex(2).String()
      pamaType, _ := obj.GetIndex(2).Int()

      pamaMessage := &netmsg.PamaMessage{
        Idkek: c.player.GetPlayerId(),
        Color: pamaColor,
        Message: pamaMsg,
        Type: pamaType,
      }

      c.broadcast <- pamaMessage.WritePacket()
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
        Idkek: c.player.GetPlayerId(),
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

      c.broadcast <- moveMessage.WritePacket()
    case "golabelnew":
      stand, _ := obj.GetIndex(1).String()
      fokStr, _ := obj.GetIndex(2).String()
      fokInt, _ := obj.GetIndex(3).Int()

      stopMessage := &netmsg.StopMessage{
        Idkek: c.player.GetPlayerId(),
        Stand: stand,
        FokStr: fokStr,
        FokInt: fokInt,
      }

      c.broadcast <- stopMessage.WritePacket()
    case "comandos":
      cmd, _ := obj.GetIndex(1).String()
      var param string = ""
      var user string = ""

      if cmd == "boing" || cmd == "rot" {
        param = strconv.FormatInt(c.player.GetPlayerId(), 10)
      } else if cmd == "alert" {
        user = c.player.GetUsername()
        param, _ = obj.GetIndex(2).String()
      }

      cmdMessage := &netmsg.CmdMessage{
        User: user,
        Cmd: cmd,
        Param: param,
      }

      c.output <- cmdMessage.WritePacket()
    default:
      fmt.Printf("Unhandled packet received - %v\n", namespace)
  }
}

func (c *Connection) Close() {
  // Close channels
	close(c.output)

  // Close the websocket
  c.ws.Close()

  c.player = nil
}
