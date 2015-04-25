package helpers

// <imports>
import (
  "fmt"
  "log"

  "github.com/eaigner/hood"
  _ "github.com/ziutek/mymysql/godrv"
)

var db *hood.Hood

func OpenDatabaseConnection() *hood.Hood {
  if db == nil {
    var username string = "root"
    var password string = "test"
    var scheme string = "kekocity"

    connectionString := fmt.Sprintf("%v/%v/%v", scheme, username, password)

    _db, err := hood.Open("mymysql", connectionString)

    if err != nil {
      log.Fatal("helpers.database:", "unable to connect to database", err.Error())
      return nil
    }

    db = _db
  }

  return db
}
