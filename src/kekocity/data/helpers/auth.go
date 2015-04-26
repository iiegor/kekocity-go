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

func (a *authHelper) userEntityToModel(_entity *entities.User) (*models.User, error) {
  u := models.NewUser(_entity, db)
	u.Username = _entity.Username

  return u, nil
}

func (a *authHelper) AuthenticateUsingCredentials(_token string) (interfaces.IUser, error) {
  var users *entities.User
  db.Query("SELECT * FROM user").Rows(&users)

  if users == nil {
		return nil, fmt.Errorf("Player '%s' not found", _token)
	}

	playerModel, err := a.userEntityToModel(users)
	if err != nil {
		return nil, err
	}

	return playerModel, nil
}
