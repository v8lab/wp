package base

import (
	"sync"
)

func NewEntry(Kind string) (Entry EntryIntf) {
	CreateFunc := GetEntryFactory()
	Entry = CreateFunc.Create(Kind)
	return
}

var SingleFactory FactoryStu
var SingleFactoryOnce sync.Once

func GetEntryFactory() *FactoryStu {
	SingleFactoryOnce.Do(SingleFactory.Init)
	return &SingleFactory
}

type CreateFuncType func() (Entry EntryIntf)

type FactoryStu struct {
	Factorys       map[string]CreateFuncType
	BaseCreateFunc CreateFuncType
}

func (r *FactoryStu) Create(Kind string) (Entry EntryIntf) {
	if CreateFunc, ok := r.Factorys[Kind]; ok {
		Entry = CreateFunc()
	} else {
		Entry = r.BaseCreateFunc()
	}
	return
}

func (r *FactoryStu) Add(Kind string, CreateFunc CreateFuncType) {
	r.Factorys[Kind] = CreateFunc
}

func (r *FactoryStu) Init() {
	r.Factorys = make(map[string]CreateFuncType)
	//	r.BaseCreateFunc = func() EntryIntf { return &EntryUnkownStu{NewEmptyEntry(KINDUNKOWN)} }
}
