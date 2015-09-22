package helpers

import (
  "log"
  "fmt"

  "github.com/eaigner/jet"
  _ "github.com/ziutek/mymysql/godrv"
)

var db *jet.Db

func OpenDatabaseConnection() *jet.Db {
  if db == nil {
    var username string = "adminUWZjs8T"
    var password string = "ItLe9K-PdEz3"
    var scheme string = "us"

    connectionString := fmt.Sprintf("%v/%v/%v", scheme, username, password)

    _db, err := jet.Open("mymysql", connectionString)

    if err != nil {
      log.Fatal("helpers.database:", "unable to connect to database", err.Error())
      return nil
    }

    db = _db
  }

  return db
}
