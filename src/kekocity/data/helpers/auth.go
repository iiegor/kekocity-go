package helpers

import (
  "fmt"

  "kekocity/data/entities"
  "kekocity/data/models"
  "kekocity/interfaces"
)

var AuthHelper *authHelper

type authHelper struct{}

func init() {
	AuthHelper = &authHelper{}
}

func (a *authHelper) playerEntityToModel(_entity *entities.Player) (*models.Player, error) {
  u := models.NewPlayer(_entity)

  return u, nil
}

func (a *authHelper) AuthenticateUsingCredentials(_token string) (interfaces.IPlayer, error) {
  var players *entities.Player
  var query string = fmt.Sprintf("SELECT * FROM user WHERE token = '%v' LIMIT 1", _token)
  db.Query(query).Rows(&players)

  if players == nil {
		return nil, fmt.Errorf("Player '%s' not found", _token)
	}

	playerModel, err := a.playerEntityToModel(players)
	if err != nil {
		return nil, err
	}

	return playerModel, nil
}
