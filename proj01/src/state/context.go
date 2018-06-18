package state

import (
	"sync"
	"time"
)

var SingleContext Context
var SingleContextOnce sync.Once

func GetSingleContext() *Context {
	SingleContextOnce.Do(SingleContext.Init)
	return &SingleContext
}

type Context struct {
	cIos        chan int
	cHuawei     chan int
	cOther      chan int
	cIosOver    chan int
	cHuaweiOver chan int
	cOtherOver  chan int
	StateIos    StateIntf
	StateHuawei StateIntf
	StateOther  StateIntf
}

func (r *Context) Init() {
	r.StateIos = GetSinleStateIos()
	r.StateHuaWei = GetSingleStateHuawei()
	r.StateOther = GetSinleStateOther()
}
func (r *Context) IosOver() {
	r.cIosOver <- 1
}
func (r *Context) HuaweiOver() {
	r.cHuaweiOver <- 1
}
func (r *Context) OtherOver() {
	r.cOtherOver <- 1
}
func (r *Context) DispatchOver() {
	for {
		select {
		case <-r.cIosOver:
			r.cHuawei <- 1
		case <-r.cHuaweiOver:
			r.cOther <- 1
		case <-r.cOtherOver:
		}
	}
}
func (r *Context) DispatchStart() {
	for {
		select {
		case <-r.cIos:
			r.StateHuawei.Stop()
			r.StateOther.Stop()
			go r.StateIos.Execute()
		default:
			select {
			case <-r.cIos:
				r.StateHuaWei.Stop()
				r.StateOther.Stop()
				go r.StateIos.Execute()
			case <-r.cHuawei:
				r.State.Stop()
				r.State = GetSingleStateHuawei()
				go r.StateHuawei.Execute()
			default:
				select {
				case <-r.cIos:
					r.StateHuaWei.Stop()
					r.StateOther.Stop()
					go r.StateIos.Execute()
				case <-r.cHuawei:
					r.StateIos.Stop()
					r.StateOther.Stop()
					go r.StateHuawei.Execute()
				case <-r.cOther:
					r.StateIos.Stop()
					r.StateHuaWei.Stop()
					go r.StateOther.Execute()
				}
			}
		}
	}
}
