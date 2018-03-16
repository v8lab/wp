package base

import (
	"sync"

	Tool "tool"
)

func NewEntryKind(Kind string) (Entry EntryIntf) {
	CreateFunc := GetSingleFactory()
	Entry = CreateFunc.Create(Kind)
	return
}

var SingleFactory EntryFactoryStu
var SingleFactoryOnce sync.Once

func GetSingleFactory() *EntryFactoryStu {
	SingleFactoryOnce.Do(SingleFactory.Init)
	return &SingleFactory
}

type CreateFunc func() (Entry EntryIntf)

type EntryFactoryStu struct {
	Factorys       map[string]CreateFunc
	UnknownFactory CreateFunc
}

func (r *EntryFactoryStu) Create(Kind string) (Entry EntryIntf) {
	if Create, ok := r.Factorys[Kind]; ok {
		Entry = Create()
	} else {
		Entry = r.UnknownFactory()
	}
	Entry.SetId(Tool.GenId())
	return
}
func (r *EntryFactoryStu) Add(Kind string, Func CreateFunc) {
	r.Factorys[Kind] = Func
}

func (r *EntryFactoryStu) Init() {
	r.Factorys = make(map[string]CreateFunc)
	r.UnknownFactory = func() EntryIntf { return &EntryUnkownStu{NewEmptyEntry(KINDUNKOWN)} }
}

type EntryUnkownStu struct {
	*EntryStu
}

func (r *EntryUnkownStu) Execute() (EntryOut *EntryStu, ret int) {

	return
}
