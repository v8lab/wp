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
	if r.Running() {
		return
	}
	for {
		select {
		case <-r.interrupt:
			r.running = false
		default:
			select {
			case data := <-r.buffer:
				if v, ok := data.(int); ok {
					go r.Task(v)
				}
			case <-r.interrupt:
				r.running = false
				return
			default:
				GetSingleContext().OtherOver()
				r.running = false
				return
			}
		}
	}
}
