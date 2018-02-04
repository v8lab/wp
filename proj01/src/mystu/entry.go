package mystu

import (
	"tool"
)

type EntryStu struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}

func (r *EntryStu) Init(Id string) {
	r.Id = Id
}

func (r *EntryStu) GetId() string {
	return r.Id
}
func (r *EntryStu) ToJson() string {
	return tool.Stu2Json(r)
}
func (r *EntryStu) InitByJson(data string) {
	tool.Json2Stu(data, r)
}
