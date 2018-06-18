package state

import (
	"sync"
)

var SingleStateIos StateIos
var SingleStateIosOnce sync.Once

func GetSinleStateIos() *StateIos {
	return &SingleStateIos
}

type StateIos struct {
	running   bool
	interrupt chan int
	buffer    []int
}

func (r *StateIos) Task(data int) {
	fmt.Println(data)
	time.Sleep(time.Second * 1)
}

func (r *StateIos) Execute() {
	for {
		select {
		case <-r.interrupt:
			r.running = false
		default:
			select {
			case data <- r.buffer:
				go r.Task(data)
			case <-r.interrupt:
				r.running = false
			default:
				GetSingleContext().IosOver()
				r.running = false
			}
		}
	}
}
func (r *StateIos) Stop() {
	if r.running {
		r.running = false
		r.interrupt <- 1
	}
}
