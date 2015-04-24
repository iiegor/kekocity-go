package kekocity

// <imports>
import (
  "log"
  "flag"
  "io"
  "kekocity/net"
  "kekocity/types"
  "net/http"
)

var addr = flag.String("addr", types.SERVICE, "http service address")

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

func Prepare() {
  log.Println("KEKOCITY-GO")
  log.Println("*****************************************")
}

func Boot() {
  flag.Parse()

  go net.Run()

  defer func() {
    // Handle functions
    http.HandleFunc("/", serveDefault)
    http.HandleFunc("/ws", net.ServeWs)

    // Start listening
    err := http.ListenAndServe(*addr, nil)

  	if err != nil {
  		log.Fatal("ListenAndServe: ", err)
  	}
  }()

  log.Println("Listening for new connections...")
}
