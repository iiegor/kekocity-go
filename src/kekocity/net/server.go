package net

// <imports>
import (
  "log"
  "fmt"
  "net/http"

  "golang.org/x/net/websocket"

  pnet "kekocity/misc/packet"
  cmap "kekocity/misc/concurrentmap"
  "kekocity/data/helpers"
  "kekocity/net/message"
)

var server *Server

type Server struct {
  port int

  connectedUsers *cmap.ConcurrentMap
}

func init() {
	server = newServer()
}

func newServer() *Server {
	return &Server{
    port: 8080,
    connectedUsers: cmap.New(),
  }
}

func Listen(_port int) {
  server.port = _port

  log.Printf("Listening for connections on port %d!", _port)

  http.Handle("/ws", websocket.Handler(clientConnection))

	err := http.ListenAndServe(fmt.Sprintf(":%d", _port), nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func clientConnection(clientsock *websocket.Conn) {
  packet := pnet.NewPacket()
  buffer := make([]uint8, pnet.PACKET_MAXSIZE)

  recv, err := clientsock.Read(buffer)

  if err == nil {
    copy(packet.Buffer[0:recv], buffer[0:recv])

    parseFirstMessage(clientsock, packet)
  } else {
    if err.Error() != "EOF" {
      log.Println("net.server", "Client connection error:", err.Error())
    }
  }
}

func parseFirstMessage(_conn *websocket.Conn, _packet *pnet.Packet) {
  _message := _packet.ToString()

  // If the first packet length is < 1 close the socket
  if len(_message) < 1 {
    _conn.Close()
    return
  }

  // Create the connection
  connection := NewConnection(_conn)

  // Authentication wrapper
  authPacket := &message.AuthMessage{}
  user, err := helpers.AuthHelper.AuthenticateUsingCredentials(_message)

  if err != nil {
    log.Fatal("Invalid credentials!")
    authPacket.Status = "error"
  } else {
    // Need to check if its already logged

    authPacket.Status = "success"

    connection.AssignToUser(user)
	  connection.txChan <- authPacket

    return
  }

  // Send bad auth message and close
  connection.txChan <- authPacket
  connection.Close()
}
