package exec

import (
	mylib "mylib"
	base "mylib/udpentry/base"
	mbase "p2p/natcomb/base"
)

var (
	KindCli01Obv01 string = base.NatStr[base.NATFull] + base.NatStr[base.NATFull]
	KindCli01Obv02 string = base.NatStr[base.NATFull] + base.NatStr[base.NATRestricted]
	KindCli01Obv03 string = base.NatStr[base.NATFull] + base.NatStr[base.NATPortRestricted]
	KindCli01Obv04 string = base.NatStr[base.NATFull] + base.NatStr[base.NATSymetric]
)

func init() {

	Factory := mbase.GetMethodFactory()
	Factory.Add(KindCli01Obv04,
		func() mbase.EntryIntf {
			return &Cli01Obv04Stu{
				EntryStu: mbase.NewEntryStu(),
			}
		})
}

type Cli01Obv04Stu struct {
	*mbase.EntryStu
	Data *base.EntryDataStu
}

func (r *Cli01Obv04Stu) Execute() (ret int) {
	mylib.PrnLog.Debug("Cli01Obv04Stu")
	r.Data = base.NewEntryDataStu(nil, nil)
	r.Data.SetKind(byte(207))
	r.Data.SetOtherAddr(r.LAddr)
	r.WriteAddr(r.Data.SData, r.RAddr)
	return
}
