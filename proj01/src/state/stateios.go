package state

import (
	"fmt"
	"sync"
	"time"
)

var SingleStateIos StateIos
var SingleStateIosOnce sync.Once

func GetSinleStateIos() *StateIos {
	return &SingleStateIos
}

type StateIos struct {
	*StateBase
}

func (r *StateIos) Task(data int) {
	fmt.Println(data)
	time.Sleep(time.Second * 1)
}

func (r *StateIos) Execute() {
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
			default:
				GetSingleContext().IosOver()
				r.running = false
			}
		}
	}
}
