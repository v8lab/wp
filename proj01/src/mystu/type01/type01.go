package mystu

import (
	"tool"
)

type EntryType01Stu struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}

func (r *EntryType01Stu) Init(Id string) {
	r.Id = Id
}

func (r *EntryType01Stu) GetId() string {
	return r.Id
}

func (r *EntryType01Stu) ToJson() string {
	return tool.Stu2Json(r)
}
func (r *EntryType01Stu) InitByJson(data string) {
	tool.Json2Stu(data, r)
}
