package entities

type Player struct {
  Id        uint
  Username  string `sql:"size(15),notnull"`
  Password  string `sql:"size(15),notnull"`
  Rank      int

  Coins     int32
  Clouds    int32

  Token     string
}
