package helpers

import (
  "database/sql"

  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func OpenDatabaseConnection() {
  var err error

  if db, err = sql.Open("mysql", "root:test@/kekocity?charset=utf8"); err != nil {
    panic(err)
  }
}
