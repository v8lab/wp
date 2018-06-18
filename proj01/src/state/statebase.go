package state

import (
	"sync"
)

type StateBase struct {
	sync.RWMutex
	running   bool
	interrupt chan int
	buffer    chan interface{}
}

func (r *StateBase) Running() bool {
	r.Lock()
	defer r.Unlock()

	if r.running {
		return true
	}

	r.running = true
	return false
}
func (r *StateBase) Stop() {
	r.Lock()
	defer r.Unlock()
	if r.running {
		r.running = false
		r.interrupt <- 1
	}
}
