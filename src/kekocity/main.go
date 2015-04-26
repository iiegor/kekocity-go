package kekocity

// <imports>
import (
  "flag"
  "io"
  "log"
  "net/http"
  "os"
  "fmt"

  "kekocity/net"
  "kekocity/types"
  "kekocity/data/helpers"
)

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
  fmt.Println("*****************************************")
  fmt.Println("**            KEKOCITY-GO              **")
  fmt.Println("*****************************************")

  if types.DEBUG {
    log.Println("Creating a database connection...")
  }
  helpers.OpenDatabaseConnection()
}

func Boot() {
  flag.Parse()

  net.Listen(types.SERVICE)
}

func Exit(code int) {
  os.Exit(code)
}
