package message

import (
  "fmt"

  "github.com/bitly/go-simplejson"
)

func Dispatch(obj *simplejson.Json) {
  fmt.Println(obj.Get("token"))
}
