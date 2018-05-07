package exec

import (
	"net"

	mylib "mylib"
	base "mylib/udpentry/base"
	clis "p2p/cli"
	method "p2p/natcomb/base"
	srv "p2p/srv"
)

var EntryObvConnReq int = 57
var EntryObvConnResp int = 58

type ObvConnRespStu struct {
	*base.EntryDataStu

	CliNatType string
	CliAddr    *net.UDPAddr
	ObvNatType string
	ObvAddr    *net.UDPAddr
}

func (r *ObvConnRespStu) DealCli() (ret int) {
	CliIdStr := r.GetCliIdStr()
	mylib.PrnLog.Debug("cli id ", CliIdStr)
	Clis := clis.GetSingleClients()
	Cli, e := Clis.Find(CliIdStr)
	if !e {
		ret = -1
		return
	}
	r.CliAddr = Cli.GetAddr()
	r.CliNatType = Cli.GetNatTypeStr()
	return
}
func (r *ObvConnRespStu) DealObv() (ret int) {
	r.ObvAddr = r.GetAddr()
	r.ObvNatType = r.GetNatTypeStr()
	return
}
func (r *ObvConnRespStu) Distribute() (ret int) {
	SrvUdp := srv.GetSingleSrvUdp()
	Factory := method.GetMethodFactory()
	Method := Factory.Create(r.CliNatType + r.ObvNatType)
	if Method == nil {
		mylib.PrnLog.Error("no method ---- ", r.CliNatType)
		mylib.PrnLog.Error("no method ---- ", r.CliAddr)
		mylib.PrnLog.Error("no method ---- ", r.ObvNatType)
		mylib.PrnLog.Error("no method ---- ", r.ObvAddr)
		return
	}
	Method.Init(r.CliAddr, r.ObvAddr, SrvUdp.GetConn())
	ret = Method.Execute()
	if ret != 0 {
		mylib.PrnLog.Error(" Method.Execute()")
	}
	return
}
func (r *ObvConnRespStu) Execute() (ret int) {
	mylib.PrnLog.Debug("Execute")
	ret = r.DealCli()
	if ret != 0 {
		mylib.PrnLog.Error("r.DealCli()")
	}
	ret = r.DealObv()
	if ret != 0 {
		mylib.PrnLog.Error("r.DealObv()")
	}
	ret = r.Distribute()
	if ret != 0 {
		mylib.PrnLog.Error("r.Distribute()")
	}
	return
}

func init() {

	ProRsv := base.GetSingleProRsvStu()
	Factory := base.GetEntryFactory()

	ProRsv.Add(EntryObvConnReq, EntryObvConnResp)

	Factory.Add(EntryObvConnResp,
		func() base.EntryIntf {
			return &ObvConnRespStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

}
