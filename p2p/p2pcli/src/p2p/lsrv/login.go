package lsrv

import (
	"sync"
	"time"

	mylib "mylib"
	base "mylib/udpentry/base"
	setting "p2p/setting"
)

var EntryCliLoginReq int = 77
var EntryCliLoginResp int = 78
var EntryCliLoginEnd int = 79

var EntryObvLoginReq int = 67
var EntryObvLoginResp int = 68
var EntryObvLoginEnd int = 69

var SingleLoginReq LoginReqStu
var SingleLoginReqOnce sync.Once

func GetSingleLoginReq() *LoginReqStu {
	SingleLoginReqOnce.Do(
		func() {
			SingleLoginReq.EntryDataStu = base.NewEntryDataStu(nil, nil)

		})
	return &SingleLoginReq
}

type LoginReqStu struct {
	*base.EntryDataStu
	Chan chan int
}

func (r *LoginReqStu) Execute() (ret int) {
	mylib.PrnLog.Debug("LoginReqStu Execute")
	SrvUdp := GetSingleSrvUdp()
	Rsrv := GetSingleRSrvInfo()
	r.Chan = make(chan int)
	SrvUdp.WriteUdp(r.SData, Rsrv.GetAddr())
	mylib.PrnLog.Debug("time")
	timer := time.NewTimer(50 * time.Second)
	select {
	case <-r.Chan:
		mylib.PrnLog.Debug("login success !")
		Setting := setting.GetSetting()
		if Setting.Model == "obv" {
			ObvConnReq := GetSingleObvConnReq()
			go ObvConnReq.Execute()
		}
	case <-timer.C:
		mylib.PrnLog.Debug("login failed !")
	}
	mylib.PrnLog.Debug("out")
	return
}

type LoginEndStu struct {
	*base.EntryDataStu
	Chan chan int
}

func (r *LoginEndStu) Execute() (ret int) {
	mylib.PrnLog.Debug(" LoginEndStu")
	SingleLoginReq.Chan <- 0
	return
}

func init() {

	ProRsv := base.GetSingleProRsvStu()
	Factory := base.GetEntryFactory()

	ProRsv.Add(EntryCliLoginReq, EntryCliLoginResp)
	ProRsv.Add(EntryCliLoginResp, EntryCliLoginEnd)

	ProRsv.Add(EntryObvLoginReq, EntryObvLoginResp)
	ProRsv.Add(EntryObvLoginResp, EntryObvLoginEnd)

	Factory.Add(EntryCliLoginReq,
		func() base.EntryIntf {
			return &LoginReqStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

	Factory.Add(EntryObvLoginReq,
		func() base.EntryIntf {
			return &LoginReqStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

	Factory.Add(EntryCliLoginEnd,
		func() base.EntryIntf {
			return &LoginEndStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

	Factory.Add(EntryObvLoginEnd,
		func() base.EntryIntf {
			return &LoginEndStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

}
