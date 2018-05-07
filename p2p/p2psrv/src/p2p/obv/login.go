package obv

import (
	mylib "mylib"
	base "mylib/udpentry/base"
)

var (
	EntryObvLogin int = 2
)

func init() {

	Factory := base.GetEntryFactory()
	Factory.Add(EntryObvLogin,
		func() base.EntryIntf {
			return &ObvLoginStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})
}

type ObvLoginStu struct {
	*base.EntryDataStu
}

func (r *ObvLoginStu) Execute() (ret int) {
	Obvs := GetSingleObservers()
	Obv := base.NewEntryDataStu(r.RData, r.Addr)
	Obvs.Add(Obv)
	mylib.PrnLog.Debug("add obv", Obv.GetId())
	return
}
