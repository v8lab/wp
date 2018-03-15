package dev

type EndPointStu struct {
	Id string `redis:"id"`
}

func (r *EndPointStu) GetId() string {
	return r.Id
}
