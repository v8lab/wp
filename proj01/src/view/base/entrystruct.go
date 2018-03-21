package base

import (
	"net/http"

	a1 "a1common"
	libf "libfunc"
	setting "view/setting"
)

func NewEmptyEntry(Kind string) *EntryStu {
	var Entry EntryStu
	Entry.Init(Kind)
	return &Entry
}

type EntryRespStu struct {
	Error int    `json:"error"`
	Msg   string `json:"msg"`
}

func (r *EntryRespStu) InitByJson(data []byte) (ret int) {
	_, ret = libf.ParseJsonByStructMsg(data, r)
	if ret != a1.Success {
		a1.PrnLog.Error("InitByJson error", string(data))
		return
	}
	return
}

type EntryStu struct {
	Id        string
	Kind      string
	Method    string
	Url       string
	Header    http.Header
	Body      string
	EntryRst  *EntryStu
	CheckFunc func() int
	Resp      EntryRespStu
}

func (r *EntryStu) Init(Kind string) {
	r.Kind = Kind
	return
}
func (r *EntryStu) Construct(unitin interface{}) (ret int) {
	unit, ok := unitin.(*setting.UnitStu)
	if !ok {
		return -1
	}
	r.SetHeader(unit.GetHeader())
	r.SetMethod(unit.GetMethod())
	r.SetUrl(unit.GetUrl())
	r.SetBody(unit.GetBody())

	return
}
func (r *EntryStu) SetKind(Kind string) {
	r.Kind = Kind
}
func (r *EntryStu) GetKind() string {
	return r.Kind
}
func (r *EntryStu) SetMethod(Method string) {
	r.Method = Method
}
func (r *EntryStu) SetUrl(Url string) {
	r.Url = Url
}
func (r *EntryStu) GetUrl() string {
	return r.Url
}
func (r *EntryStu) GetMethod() string {
	return r.Method
}
func (r *EntryStu) SetHeader(Header http.Header) {
	r.Header = Header
}
func (r *EntryStu) GetHeader() http.Header {
	return r.Header
}
func (r *EntryStu) SetBody(Body string) {
	r.Body = Body
}
func (r *EntryStu) GetBody() string {
	return r.Body
}
func (r *EntryStu) GetThis() *EntryStu {
	return r
}
func (r *EntryStu) SetCheckFunc(CheckFunc func() int) {
	r.CheckFunc = CheckFunc
}
func (r *EntryStu) ToJson() string {
	return libf.SructToJsonStringOne(r)
}

func (r *EntryStu) TransLocal() (ret int) {
	Trans := NewEntryLocalStu(r.GetThis())
	r.EntryRst, ret = Trans.Execute()
	if ret != 0 {
		a1.PrnLog.Error("Trans.Execute")
		return
	}
	return
}

func (r *EntryStu) CheckRst() (ret int) {
	ret = r.Resp.InitByJson([]byte(r.EntryRst.GetBody()))
	if ret != 0 {
		a1.PrnLog.Error("r.Resp.InitByJson")
		return
	}

	ret = r.Resp.Error
	if ret != 0 {
		return
	}
	return
}

func (r *EntryStu) Execute() (EntryOut *EntryStu, ret int) {
	defer func() {
		EntryOut = r.EntryRst

	}()

	ret = r.TransLocal()
	if ret != 0 {
		a1.PrnLog.Error("r.TransLocal")
		return
	}
	ret = r.CheckRst()
	if ret != 0 {
		return
	}
	return
}
