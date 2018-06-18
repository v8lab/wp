package state

import (
	"fmt"
	"sync"
	"time"
)

var SingleStateHuawei StateHuawei
var SingleStateHuaweiOnce sync.Once

func GetSingleStateHuawei() *StateHuawei {
	return &SingleStateHuawei
}

type StateHuawei struct {
	sync.RWMutex
	running   bool
	interrupt chan int
	buffer    []int
}

func (r *StateHuawei) Task(data int) {
	fmt.Println(data)
	time.Sleep(time.Second * 1)
}

func (r *StateHuawei) Execute() {
	if r.Running() {
		return
	}
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
				GetSingleContext().HuaweiOver()
				r.running = false
			}
		}
	}
}
func (r *StateHuawei) Running() bool {
	r.Lock()
	defer r.Unlock()

	if r.running {
		return true
	}

	r.running = true
	return false
}
func (r *StateHuawei) Stop() {
	r.Lock()
	defer r.Unlock()
	if r.running() {
		r.running = false
		r.interrupt <- 1
	}
}
