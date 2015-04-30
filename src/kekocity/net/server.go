package net

import (
  "fmt"
  "net/http"

  "github.com/gorilla/websocket"

  "kekocity/data/helpers"
  netmsg "kekocity/net/message"
)

var origins = map[string]bool {
  "http://localhost": true, // development
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
  if !origins[r.Header.Get("Origin")] {
		http.Error(w, "Origin not allowed", 403)
		return
	}
  //Handshake
  ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
  if _, ok := err.(websocket.HandshakeError); ok {
    http.Error(w, "Not a websocket handshake", 400)
    return
  } else if err != nil {
    return
	}

  // Get possible session params
  playerId := r.URL.Query().Get("kakid")
  playerToken := r.URL.Query().Get("kakpa")

  if len(playerId) == 0 || len(playerToken) == 0 {
    ws.Close()
    return
  }

  // Add connection
  connection := NewConnection(ws)

  authMessage := &netmsg.AuthMessage{}
  player, err := helpers.AuthHelper.AuthenticateUsingCredentials(playerToken)
  if err != nil {
    authMessage.Status = "bad_credentials"
  } else {
    authMessage.Status = "success"

    connection.AssignToPlayer(player)
    connection.output <- authMessage.WritePacket()

    return
  }

  connection.output <- authMessage.WritePacket()
  connection.Close()
}

func Listen(_service int) {
  http.HandleFunc("/ws", wsHandler)

  fmt.Printf("Listening for connections on %d!\n", _service)

  if err := http.ListenAndServe(fmt.Sprintf(":%d", _service), nil); err != nil {
		panic(err)
	}
}
