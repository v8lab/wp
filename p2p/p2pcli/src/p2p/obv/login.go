package obv

import (
	mylib "mylib"
	base "mylib/udpentry/base"
	"net"
	lsrv "p2p/lsrv"
	"time"
)

var (
	EntryOtherLoginReq  int = 207
	EntryOtherLoginResp int = 208
	EntryOtherLoginEnd  int = 209
)

type EntryOtherLoginRespStu struct {
	*base.EntryDataStu
}

func (r *EntryOtherLoginRespStu) Execute() (ret int) {
	mylib.PrnLog.Debug(" EntryOtherLoginRespStu")
	SrvUdp := lsrv.GetSingleSrvUdp()
	OtherAddr := r.GetOtherAddr()
	mylib.PrnLog.Debug(" OtherAddr", OtherAddr)
	r.SetKind(byte(EntryOtherLoginResp))
	r.SetId(12349999)
	SrvUdp.WriteUdp(r.SData, OtherAddr)
	go r.Heart(OtherAddr)
	return
}

func (r *EntryOtherLoginRespStu) Heart(Addr *net.UDPAddr) (ret int) {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		SrvUdp := lsrv.GetSingleSrvUdp()
		r.SetKind(byte(EntryOtherLoginResp))
		r.SetId(12349999)
		SrvUdp.WriteUdp(r.SData, Addr)
	}
	return
}

type EntryOtherLoginEndStu struct {
	*base.EntryDataStu
}

func (r *EntryOtherLoginEndStu) Execute() (ret int) {
	mylib.PrnLog.Debug(" EntryOtherLoginEndStu")
	Obvs := GetSingleObservers()
	Obv := base.NewEntryDataStu(r.RData, r.Addr)
	Obvs.Add(Obv)
	mylib.PrnLog.Debug("cli Add obvs", Obv.GetId())
	return
}

func init() {

	ProRsv := base.GetSingleProRsvStu()
	Factory := base.GetEntryFactory()

	ProRsv.Add(EntryOtherLoginReq, EntryOtherLoginResp)
	ProRsv.Add(EntryOtherLoginResp, EntryOtherLoginEnd)

	Factory.Add(EntryOtherLoginResp,
		func() base.EntryIntf {
			return &EntryOtherLoginRespStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

	Factory.Add(EntryOtherLoginEnd,
		func() base.EntryIntf {
			return &EntryOtherLoginEndStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

}
