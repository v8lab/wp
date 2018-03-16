package sysinfo

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	. "base"
	client "client"
)

const (
	KindCliList string = "/sysinfo/clilist"
)

func init() {
	Factory := GetSingleFactory()
	Factory.Add(KindCliList, func() EntryIntf { return &CliListStu{EntryStu: NewEmptyEntry(KindCliList)} })

	ProRso := GetSingleProblemResolveMap()
	ProRso.Add(KindCliList, KindCliList)
}

type CliListStu struct {
	*EntryStu
	Format   string
	Tw       *tabwriter.Writer
	Buf      *bytes.Buffer
	EntryRst *EntryStu
}

func (r *CliListStu) InitPrn() (ret int) {
	r.Format = "%v\t%v\t\n"
	r.Buf = new(bytes.Buffer)
	r.Tw = new(tabwriter.Writer).Init(r.Buf, 0, 8, 2, ' ', 0)
	r.PrnLog("clients list", " detail ")
	return
}
func (r *CliListStu) PrnLog(arg1, arg2 string) {
	fmt.Fprintf(r.Tw, r.Format, arg1, arg2)
}
func (r *CliListStu) ScanClients() (ret int) {
	Clients := client.GetClientsStu()
	for _, v := range Clients.Clients {
		r.PrnLog("", v.GetId())
	}
	return
}
func (r *CliListStu) CreateRst() (ret int) {
	r.EntryRst = NewEmptyEntry("")
	r.Tw.Flush()
	r.EntryRst.SetBody(r.Buf.String())
	return
}
func (r *CliListStu) Execute() (EntryOut *EntryStu, ret int) {
	defer func() {
		EntryOut = r.EntryRst
	}()
	ret = r.InitPrn()
	if ret != 0 {
		r.PrnLog("error", "r.InitPrn")
		return
	}
	ret = r.ScanClients()
	if ret != 0 {
		r.PrnLog("error", "r.ScanClients")
		return
	}
	ret = r.CreateRst()
	if ret != 0 {
		r.PrnLog("error", "r.CreateRst")
		return
	}
	return
}
