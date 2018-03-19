package sysinfo

import (
	. "base"
	client "client"
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

func (r *ClientStu) GetInfoClis() (ret int) {
	r.Print("client", "--list--")
	r.Print("-----", "----------")
	Clients := client.GetClientsStu()
	for k, v := range Clients.Clients {
		r.Print(k, v.GetId())
	}
	return
}
func (r *ClientStu) GetInfoCli() (ret int) {
	r.Print("client", r.ClientId)
	r.Print("-----", "----------")
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
func (r *ClientStu) GetInfo() (ret int) {
	r.ClientId = r.Form.Get("id")
	if len(r.ClientId) == 0 {
		ret = r.GetInfoClis()
		if ret != 0 {
			r.Print("error", "r.ScanClients")
			return
		}
	} else {
		ret = r.GetInfoCli()
		if ret != 0 {
			r.Print("error", "r.ScanClients")
			return
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

	ret = r.InitPrn()
	if ret != 0 {
		r.Print("error", "r.InitPrn")
		return
	}
	ret = r.GetInfo()
	if ret != 0 {
		r.Print("error", "r.GetInfo")
		return
	}
	ret = r.CreateRst()
	if ret != 0 {
		r.Print("error", "r.CreateRst")
		return
	}
	return
}
