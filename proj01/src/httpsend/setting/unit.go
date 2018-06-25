package setting

import (
	"encoding/xml"
	"fmt"
	"net/http"

	libf "libfunc"
)

type UnitBodyStu struct {
	XMLName xml.Name `xml:"body"       json:"-"`
	Body    string   `xml:"body"       json:"body"`
}

func (r *UnitBodyStu) Init() (ret int) {

	return
}

func (r *UnitBodyStu) ToJson() string {
	return libf.SructToJsonStringOne(r)
}

type UnitStu struct {
	XMLName    xml.Name    `xml:"unit"   json:"-"`
	Kind       string      `xml:"kind"   json:"kind"`
	Step       string      `xml:"step"   json:"step"`
	Method     string      `xml:"method" json:"method"`
	Url        string      `xml:"url"    json:"url"`
	HeaderData string      `xml:"header" json:"-"`
	Header     http.Header `xml:"-"      json:"header"`
	BodyData   string      `xml:"body"   json:"body"`
}

func (r *UnitStu) Init() (ret int) {
	ret = r.InitHeader()
	if ret != 0 {
		return
	}
	return
}

func (r *UnitStu) InitHeader() (ret int) {
	_, ret = libf.ParseJsonByStructMsg([]byte(r.HeaderData), r)
	if ret != 0 {
		fmt.Println("error ParseJsonByStructMsg ")
		return
	}
	return
}
func (r *UnitStu) GetData() interface{} {
	return r
}
func (r *UnitStu) Update() (ret int) {
	return
}
func (r *UnitStu) GetStep() string {
	return r.Step
}
func (r *UnitStu) GetKind() string {
	return r.Kind
}
func (r *UnitStu) GetId() string {
	return r.Step
}
func (r *UnitStu) GetMethod() string {
	return r.Method
}
func (r *UnitStu) GetUrl() string {
	return r.Url
}
func (r *UnitStu) GetHeader() http.Header {
	return r.Header
}
func (r *UnitStu) GetBody() string {
	return r.BodyData
}

func (r *UnitStu) PrintFailed() {
	fmt.Println(r.GetStep(), r.Kind, " ---- error ")
	fmt.Println("\tMethod      --> ", r.GetMethod())
	fmt.Println("\tUrl         --> ", r.GetUrl())
	fmt.Println("\tHeader      --> ", r.GetHeader())
	fmt.Println("\tBody        --> ", r.GetBody())

}
func (r *UnitStu) PrintSuccess() {
	fmt.Println(r.GetStep(), r.Kind, " ---- success")
}
