package client

func NewClientStu(Id string) *ClientStu {
	var cli ClientStu
	cli.Init(Id)
	return &cli
}

type ClientStu struct {
	Id string
}

func (r *ClientStu) Init(Id string) {
	r.Id = Id
}

func (r *ClientStu) GetThis() interface{} {
	return r
}
