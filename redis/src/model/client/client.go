package client

import (
	dev "model/client/dev"
)

type ClientStu struct {
	Id    string        `redis:"id"`
	Rooms []dev.RoomStu `redis:"-"`
}

func (r *ClientStu) GetId() string {
	return r.Id
}
