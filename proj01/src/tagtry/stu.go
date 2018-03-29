package tagtry

type ClientStu struct {
	Id   string  `json:"id"`
	Data string  `json:"data"`
	Room RoomStu `json:"room"`
}

func (r *ClientStu) Execute() (ret int) {
	return
}

type RoomStu struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}

func (r *RoomStu) Execute() (ret int) {
	return
}
