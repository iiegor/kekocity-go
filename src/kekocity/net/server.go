package net

import (
  "fmt"
  "net/http"

  "github.com/gorilla/websocket"
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
  //Handshake
  ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
  if _, ok := err.(websocket.HandshakeError); ok {
    http.Error(w, "Not a websocket handshake", 400)
    return
  } else if err != nil {
    return
	}

  // Add connection
  NewConnection(ws)
}

func Listen(_service int) {
  http.HandleFunc("/ws", wsHandler)

  fmt.Printf("Listening for connections on %d!\n", _service)

  if err := http.ListenAndServe(fmt.Sprintf(":%d", _service), nil); err != nil {
		panic(err)
	}
}
