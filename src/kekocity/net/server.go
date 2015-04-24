package net

// <imports>
import (
  "log"
  "fmt"
  "net/http"

  "golang.org/x/net/websocket"
  pnet "kekocity/net/packet"
)

func clientConnection(clientsock *websocket.Conn) {
  log.Println("New client connection!")

  packet := pnet.NewPacket()
  buffer := make([]uint8, pnet.PACKET_MAXSIZE)

  recv, err := clientsock.Read(buffer)

  if err == nil {
    copy(packet.Buffer[0:recv], buffer[0:recv])

    var p string = packet.ToString()

    log.Println("Message received:", p)
  } else {
    if err.Error() != "EOF" {
      log.Println("GameServer", "clientConnection", "Client connection error: %v", err.Error())
    }
  }
}

func Listen(_port int) {
  log.Println("Listening for new connections...")

  http.Handle("/ws", websocket.Handler(clientConnection))

	err := http.ListenAndServe(fmt.Sprintf(":%d", _port), nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
