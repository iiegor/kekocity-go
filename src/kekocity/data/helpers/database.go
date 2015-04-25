package helpers

// <imports>
import (
  "fmt"
  "log"

  "github.com/eaigner/hood"
  _ "github.com/ziutek/mymysql/godrv"
)

func OpenDatabaseConnection() *hood.Hood {
  var username string = "root"
	var password string = "test"
	var scheme string = "kekocity"
  
	connectionString := fmt.Sprintf("%v/%v/%v", scheme, username, password)

  hd, err := hood.Open("mymysql", connectionString)

  if err != nil {
    log.Fatal("helpers.database:", "unable to connect to database", err.Error())
    return nil
  }

  return hd
}
