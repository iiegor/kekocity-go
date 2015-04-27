package message

import (
  "github.com/bitly/go-simplejson"

  "kekocity/interfaces"
  "kekocity/data/helpers"
)

func AuthPacket(_data *simplejson.Json) (interfaces.IPlayer, error) {
  token, err := _data.Get("token").String()

  if len(token) < 1 || err != nil {
    return nil, err
  }

  return helpers.AuthHelper.AuthenticateUsingCredentials(token)
}
