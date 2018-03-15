package base

import (
	"sync"

	Tool "tool"
)

func NewEntryKind(Kind string) (Entry EntryIntf) {
	CreateFunc := GetSingleEntryFactory()
	Entry = CreateFunc.Create(Kind)
	return
}

var SingleEntryFactory EntryFactoryStu
var SingleEntryFactoryOnce sync.Once

func GetSingleEntryFactory() *EntryFactoryStu {
	SingleEntryFactoryOnce.Do(SingleEntryFactory.Init)
	return &SingleEntryFactory
}

type EntryCreateFunc func() (Entry EntryIntf)

type EntryFactoryStu struct {
	Factorys       map[string]EntryCreateFunc
	UnknownFactory EntryCreateFunc
}

func (r *EntryFactoryStu) Create(Kind string) (Entry EntryIntf) {
	if EntryCreateFunc, ok := r.Factorys[Kind]; ok {
		Entry = EntryCreateFunc()
	} else {
		Entry = r.UnknownFactory()
	}
	Entry.SetId(Tool.GenId())
	return
}
func (r *EntryFactoryStu) Add(Kind string, CreateFunc EntryCreateFunc) {
	r.Factorys[Kind] = CreateFunc
}

func (r *EntryFactoryStu) Init() {
	r.Factorys = make(map[string]EntryCreateFunc)
	r.UnknownFactory = func() EntryIntf { return &EntryUnkownStu{NewEmptyEntry(KINDUNKOWN)} }
}

type EntryUnkownStu struct {
	*EntryStu
}

func (r *EntryUnkownStu) Execute() (ret int) {

	return
}
