package exec

import (
	mylib "mylib"
	base "mylib/udpentry/base"
	mbase "p2p/natcomb/base"
)

var (
	KindCli02Obv02 string = base.NatStr[base.NATRestricted] + base.NatStr[base.NATRestricted]
	KindCli02Obv03 string = base.NatStr[base.NATRestricted] + base.NatStr[base.NATPortRestricted]
	KindCli02Obv04 string = base.NatStr[base.NATRestricted] + base.NatStr[base.NATSymetric]

	KindCli03Obv02 string = base.NatStr[base.NATPortRestricted] + base.NatStr[base.NATRestricted]
	KindCli03Obv03 string = base.NatStr[base.NATPortRestricted] + base.NatStr[base.NATPortRestricted]

	KindCli04Obv02 string = base.NatStr[base.NATSymetric] + base.NatStr[base.NATRestricted]
)

func init() {

	Factory := mbase.GetMethodFactory()
	Factory.Add(KindCli02Obv02,
		func() mbase.EntryIntf {
			return &Cli02Obv02Stu{
				EntryStu: mbase.NewEntryStu(),
			}
		})
}

type Cli02Obv02Stu struct {
	*mbase.EntryStu
	Data01 *base.EntryDataStu
	Data02 *base.EntryDataStu
}

func (r *Cli02Obv02Stu) Execute() (ret int) {
	mylib.PrnLog.Debug("Cli02Obv02Stu")

	r.Data01 = base.NewEntryDataStu(nil, nil)
	r.Data01.SetKind(byte(207))
	r.Data01.SetOtherAddr(r.LAddr)
	r.WriteAddr(r.Data01.SData, r.RAddr)

	r.Data01 = base.NewEntryDataStu(nil, nil)
	r.Data01.SetKind(byte(207))
	r.Data01.SetOtherAddr(r.RAddr)
	r.WriteAddr(r.Data01.SData, r.LAddr)

	return
}
