package client

func NewClientStu(Id string) *ClientStu {
	var cli ClientStu
	cli.Init(Id)
	return &cli
}

type ClientStu struct {
	Id   string
	Data string
}

func (r *ClientStu) Init(Id string) {
	r.Id = Id
}

func (r *ClientStu) GetThis() interface{} {
	return r
}
func (r *ClientStu) GetData() interface{} {
	return r
}
func (r *ClientStu) Update() (ret int) {
	return
}
func (r *ClientStu) GetId() (Id string) {
	return r.Id
}

func (r *ClientStu) GetDetail() (kv map[string]string) {
	kv = make(map[string]string, 0)
	kv["id"] = r.Id
	kv["data"] = r.Data
	return
}
