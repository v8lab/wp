package client

type ClientsStu struct {
	Id         string                 `redis:"id"`
	Clients    []ClientStu            `redis:"-"`
	ClientsStr map[string]interface{} `redis:"clients"`
}

func (r *ClientsStu) GetId() string {
	return r.Id
}
