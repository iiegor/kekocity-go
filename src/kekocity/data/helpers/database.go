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
    var username string = "root"
    var password string = "test"
    var scheme string = "kekocity"

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
