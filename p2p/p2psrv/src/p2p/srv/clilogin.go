package srv

import (
	base "mylib/udpentry/base"

	mylib "mylib"
	clis "p2p/cli"
)

var EntryCliLoginReq int = 77
var EntryCliLoginResp int = 78
var EntryCliLoginEnd int = 79

type CliLoginRespStu struct {
	*base.EntryDataStu
	Chan chan int
}

func (r *CliLoginRespStu) Execute() (ret int) {
	SrvUdp := GetSingleSrvUdp()
	r.SetKind(byte(EntryCliLoginResp))
	mylib.PrnLog.Debug("r.String()", r.String())
	SrvUdp.WriteAddr(r.SData, r.Addr)

	Clis := clis.GetSingleClients()
	Cli := base.NewEntryDataStu(r.RData, r.Addr)
	Clis.Add(Cli)
	mylib.PrnLog.Debug("add cli", Cli.GetId())
	return
}

func init() {

	ProRsv := base.GetSingleProRsvStu()
	Factory := base.GetEntryFactory()

	ProRsv.Add(EntryCliLoginReq, EntryCliLoginResp)
	ProRsv.Add(EntryCliLoginResp, EntryCliLoginEnd)

	Factory.Add(EntryCliLoginResp,
		func() base.EntryIntf {
			return &CliLoginRespStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

}
