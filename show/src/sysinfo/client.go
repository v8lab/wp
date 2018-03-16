package sysinfo

import (
	. "base"
	client "client"
	"fmt"
	tool "tool"
)

const (
	KindClient string = "/sysinfo/client"
)

func init() {
	Factory := GetSingleFactory()
	Factory.Add(KindClient, func() EntryIntf { return &ClientStu{EntryStu: NewEmptyEntry(KindClient)} })

	ProRso := GetSingleProblemResolveMap()
	ProRso.Add(KindClient, KindClient)
}

type ClientStu struct {
	*EntryStu
	EntryRst *EntryStu
	*tool.TwLogStu
	ClientId string
}

func (r *ClientStu) InitPrn() (ret int) {
	r.TwLogStu = tool.NewTwLogStu("%v\t%v\t\n")
	return
}

func (r *ClientStu) ScanClient() (ret int) {
	Clients := client.GetClientsStu()
	Client := Clients.Find(r.ClientId)
	if Client != nil {
		kv := Client.GetDetail()
		for k, v := range kv {
			r.Print(k, v)
		}
	}
	return
}
func (r *ClientStu) CreateRst() (ret int) {
	r.EntryRst = NewEmptyEntry("")
	r.EntryRst.SetBody(r.Flush())
	return
}
func (r *ClientStu) Execute() (EntryOut *EntryStu, ret int) {
	defer func() {
		EntryOut = r.EntryRst
	}()
	fmt.Println("enter ")
	ret = r.InitPrn()
	r.ClientId = r.Form.Get("id")
	r.Print("client", r.ClientId)
	if ret != 0 {
		r.Print("error", "r.InitPrn")
		return
	}
	ret = r.ScanClient()
	if ret != 0 {
		r.Print("error", "r.ScanClients")
		return
	}
	ret = r.CreateRst()
	if ret != 0 {
		r.Print("error", "r.CreateRst")
		return
	}
	return
}
