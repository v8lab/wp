package base

import (
	libf "libfunc"
)

var (
	KINDLOCAL string = "kindLocal"
)

func init() {
	Factory := GetEntryFactory()
	Factory.Add(KINDLOCAL, func() EntryIntf { return &EntryLocalStu{EntryStu: NewEmptyEntry(KINDLOCAL)} })

}
func NewEntryLocalStu(Entry *EntryStu) EntryIntf {
	EntryLocal := EntryLocalStu{EntryStu: NewEmptyEntry(KINDLOCAL)}
	EntryLocal.InitByEntry(Entry)
	return &EntryLocal
}

type EntryLocalStu struct {
	*EntryStu
}

func (r *EntryLocalStu) InitByEntry(Entry *EntryStu) {
	r.SetMethod(Entry.GetMethod())
	r.SetUrl(Entry.GetUrl())
	r.SetHeader(Entry.GetHeader())
	r.SetBody(Entry.GetBody())
}
func (r *EntryLocalStu) Execute() (EntryOut *EntryStu, ret int) {
	EntryOut = NewEmptyEntry("response")
	Body, _, Header, _, _, ret := libf.SendHttpOriginal(r.GetMethod(), "https://nanjing.ibroadlink.com:8443"+r.GetUrl(), []byte(r.GetBody()), r.GetHeader())
	if ret == 0 {
		EntryOut.SetBody(string(Body))
		EntryOut.SetHeader(Header)
	}
	return
}
func (r *EntryLocalStu) Construct(interface{}) (ret int) {
	return
}
