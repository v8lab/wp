package base

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	KINDUNKOWN string = "UNKNOWN"
)

func NewEmptyEntry(Kind string) *EntryStu {
	var Entry EntryStu
	Entry.Kind = Kind
	return &Entry
}

type EntryStu struct {
	Id         string      `json:"msgid"`
	PathId     string      `json:"pathid"`
	Kind       string      `json:"framekind"`
	ExecStatus int         `json:"execstatus"`
	ErrCode    int         `json:"errcode"`
	Method     string      `json:"method"`
	Url        string      `json:"url"`
	Form       url.Values  `json:"url"`
	Header     http.Header `json:"headers"`
	Body       string      `json:"body"`
}

func (r *EntryStu) Init(req *http.Request) (ret int) {

	Mothod := req.Method

	r.SetMethod(Mothod)

	Url := fmt.Sprintf("%v", req.URL)
	r.SetUrl(Url)

	r.SetHeader(req.Header)
	err := req.ParseForm()
	if err != nil {
		return -1
	}
	r.Form = req.Form

	Body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return -1
	}
	r.SetBody(string(Body))

	return
}
func (r *EntryStu) SetId(Id string) {
	r.Id = Id
}
func (r *EntryStu) GetId() (Id string) {
	return r.Id
}
func (r *EntryStu) SetKind(Kind string) {
	r.Kind = Kind
}
func (r *EntryStu) GetKind() (Kind string) {
	return r.Kind
}
func (r *EntryStu) SetMethod(Method string) {
	r.Method = Method
}
func (r *EntryStu) GetMethod() (Method string) {
	return r.Method
}
func (r *EntryStu) SetUrl(Url string) {
	r.Url = Url
}
func (r *EntryStu) GetUrl() (Url string) {
	return r.Url
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
func (r *EntryStu) GetBody() (Body string) {
	return r.Body
}
func (r *EntryStu) ToJson() string {
	Data, _ := json.Marshal(r)
	return string(Data)
}

func (r *EntryStu) GetData() *EntryStu {
	return r
}
func (r *EntryStu) Nop() {
}
