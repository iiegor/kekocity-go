package kekocity

import (
  "fmt"

  "kekocity/net"
  "kekocity/data/helpers"
  . "kekocity/types"
)

func Prepare() {
  // Print header
  fmt.Println("*****************************************")
  fmt.Println("**            KEKOCITY-GO              **")
  fmt.Println("*****************************************")

  if DEBUG {
    fmt.Println("Creating a database connection...")
  }
  helpers.OpenDatabaseConnection()
}

func Boot() {
  net.Listen(SERVICE)
}
