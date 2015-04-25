package entities

import "github.com/eaigner/hood"

type User struct {
  Id        hood.Id
  Username  string `sql:"size(15),notnull"`
  Password  string `sql:"size(15),notnull"`
  Rank      int

  Coins     int32
  Clouds    int32

  Token     string
}
