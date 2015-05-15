package helpers

import (
  "fmt"

  "kekocity/data/entities"
  "kekocity/data/models"
  "kekocity/interfaces"
)

var RoomHelper *roomHelper

type roomHelper struct{}

func init() {
	RoomHelper = &roomHelper{}
}

func (rh *roomHelper) roomEntityToModel(_entity *entities.Room) (*models.Room, error) {
  u := models.NewRoom(_entity)

  return u, nil
}

func (rh *roomHelper) EnterRoom(_id string) (interfaces.IRoom, error) {
  var room *entities.Room
  var query string = fmt.Sprintf("SELECT * FROM rooms WHERE id = '%v' LIMIT 1", _id)
  db.Query(query).Rows(&room)

  if room == nil {
		return nil, fmt.Errorf("Room '%s' not found", _id)
	}

	roomModel, err := rh.roomEntityToModel(room)
	if err != nil {
		return nil, err
	}

	return roomModel, nil
}
