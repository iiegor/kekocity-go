package kekocity

// <imports>
import (
  "log"
  "flag"
  "io"
  "net/http"
  "kekocity/types"
  "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
  ReadBufferSize:  1024,
	WriteBufferSize: 1024,
  CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var addr = flag.String("addr", types.SERVICE, "http service address")

type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

func serveDefault(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found.", 404)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<html><body><h1>It works!</h1></body></html>")
}

func Boot() {
  flag.Parse()

  defer func() {
    // Handle functions
    http.HandleFunc("/", serveDefault)

    // Start listening
    err := http.ListenAndServe(*addr, nil)

  	if err != nil {
  		log.Fatal("ListenAndServe: ", err)
  	}
  }()

  log.Println("Listening for new connections...")
}
