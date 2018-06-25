package setting

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"

	. "httpsend/intfscan"
)

func NewSettingFile(filePath string) *SettingStu {
	var file SettingStu
	file.Init(filePath)
	return &file
}

type SettingStu struct {
	XMLName  xml.Name      `xml:"setting"`
	Id       string        `xml:"settingid"`
	Cases    []*CaseStu    `xml:"case"`
	Aggre    AggregateIntf `xml:"-"`
	FilePath string        `xml:"-"`
}

func (r *SettingStu) Init(filePath string) {
	r.FilePath = filePath
	ret := r.ReadFile()
	if ret != 0 {
		fmt.Println("readfile error")
		return
	}

	r.Aggre = NewAggregateStu()
	for _, v := range r.Cases {
		vv := v
		vv.Init()
		r.Aggre.Add(vv.GetId(), vv)
	}
}

func (r *SettingStu) ReadFile() (ret int) {
	content, err := ioutil.ReadFile(r.FilePath)
	if err != nil {
		log.Fatal(err)
		ret = -1
		return
	}

	err = xml.Unmarshal(content, r)
	if err != nil {
		fmt.Println("ReadFile error", err)
		ret = -1
	}
	return
}
func (r *SettingStu) GetData() interface{} {
	return r
}

func (r *SettingStu) Update() (ret int) {
	return
}

func (r *SettingStu) GetId() string {
	return r.Id
}
func (r *SettingStu) Find(Id string) (o *CaseStu, exist bool) {
	v, e := r.Aggre.Find(Id)
	if e {
		if vv, ok := v.(*CaseStu); ok {
			o = vv
			exist = true
		}
	}
	return
}

func (r *SettingStu) Iterator() IterIntf {
	return r.Aggre.Iterator()
}
