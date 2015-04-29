package message

import (
  "fmt"

  "github.com/bitly/go-simplejson"
)

func MakeNilMessage(namespace string) *simplejson.Json {
  bytes := []byte(fmt.Sprintf(`["%v", {}]`, namespace))
  json, _ := simplejson.NewJson(bytes)

  return json
}

func MakeMessage(namespace string, json *simplejson.Json) *simplejson.Json {
  str, err := json.Encode()
  if err != nil {
    return MakeNilMessage(namespace)
  }

  bytes := []byte(fmt.Sprintf(`["%v", %v]`, namespace, string(str)))
  message, err := simplejson.NewJson(bytes)
  if err != nil {
    return MakeNilMessage(namespace)
  }

  return message
}
