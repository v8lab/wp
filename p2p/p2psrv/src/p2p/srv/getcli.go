package srv

import (
	mylib "mylib"
	base "mylib/udpentry/base"

	clis "p2p/cli"
)

var (
	EntryGetDev int = 12
)

func init() {

	Factory := base.GetEntryFactory()
	Factory.Add(EntryGetDev,
		func() base.EntryIntf {
			return &GetDevStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})
}

type GetDevStu struct {
	*base.EntryDataStu
}

func (r *GetDevStu) Execute() (ret int) {
	Clis := clis.GetSingleClients()
	Cli := base.NewEntryDataStu(r.RData, r.Addr)
	Clis.Add(Cli)
	mylib.PrnLog.Debug("add cli", Cli.GetId())
	return
}
