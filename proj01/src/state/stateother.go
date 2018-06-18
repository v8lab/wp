package state

import (
	"fmt"
	"sync"
	"time"
)

var SingleStateOther StateOther
var SingleStateOtherOnce sync.Once

func GetSinleStateOther() *StateOther {
	return &SingleStateOther
}

type StateOther struct {
	*StateBase
}

func (r *StateOther) Task(data int) {
	fmt.Println(data)
	time.Sleep(time.Second * 1)
}

func (r *StateOther) Execute() {

}
func (r *StateOther) Stop() {
	if r.running {
		r.running = false
		r.interrupt <- 1
	}
}
