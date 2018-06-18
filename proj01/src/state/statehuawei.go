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
	*StateBase
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
			case data := <-r.buffer:
				if v, ok := data.(int); ok {
					go r.Task(v)
				}
			case <-r.interrupt:
				r.running = false
			default:
				GetSingleContext().HuaweiOver()
				r.running = false
			}
		}
	}
}
