package cli

import (
	mylib "mylib"
	base "mylib/udpentry/base"
)

var (
	EntryCliLogin int = 1
)

func init() {

	Factory := base.GetEntryFactory()
	Factory.Add(EntryCliLogin,
		func() base.EntryIntf {
			return &CliLoginStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})
}

type CliLoginStu struct {
	*base.EntryDataStu
}

func (r *CliLoginStu) Execute() (ret int) {
	Clis := GetSingleClients()
	Cli := base.NewEntryDataStu(r.RData, r.Addr)
	Clis.Add(Cli)
	mylib.PrnLog.Debug("add cli", Cli.GetId())
	return
}
