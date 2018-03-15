package dev

type RoomStu struct {
	Id        string        `redis:"id"`
	EndPoints []EndPointStu `redis:"-"`
}

func (r *RoomStu) GetId() string {
	return r.Id
}
