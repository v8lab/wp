package client

import (
	"sync"
)

var SingleClients ClientsStu
var SingleClientsOnce sync.Once

func GetClientsStu() *ClientsStu {
	SingleClientsOnce.Do(SingleClients.Init)
	return &SingleClients
}

type ClientsStu struct {
	Id      string
	Clients []*ClientStu
}

func (r *ClientsStu) Init() {
	r.Clients = make([]*ClientStu, 0)
}
func (r *ClientsStu) Add(v *ClientStu) {
	r.Clients = append(r.Clients, v)
}
func (r *ClientsStu) Find(Id string) (Client *ClientStu) {
	for _, v := range r.Clients {
		if v.GetId() == Id {
			Client = v
			return
		}
	}
	return
}

var ClientId01 string = "clientId01"
var ClientId02 string = "clientId02"

func init() {
	Clients := GetClientsStu()
	Client01 := NewClientStu(ClientId01)
	Client01.Data = "qwertyuio"
	Client02 := NewClientStu(ClientId02)
	Client02.Data = "asfghjkl"
	Clients.Add(Client01)
	Clients.Add(Client02)
}
