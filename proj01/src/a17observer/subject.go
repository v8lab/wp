package a17observer

import (
	"sync"
)

type SubjectStu struct {
	Observers []ObserverIntf
	sync.RWMutex
}

func (r *SubjectStu) Init() {

}
func (r *SubjectStu) Register(v ObserverIntf) {
	r.Lock()
	defer r.Unlock()
	r.Observers = append(r.Observers, v)
}
func (r *SubjectStu) UnRegister(v ObserverIntf) {

}
func (r *SubjectStu) Notify() {
	for _, v := range r.Observers {
		v.Update()
	}
}
func (r *SubjectStu) Change() {
	r.Notify()
}
