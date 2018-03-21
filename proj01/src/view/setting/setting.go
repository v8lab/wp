package setting

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"sync"

	a1 "a1common"
	. "a1common/intfscan"
)

var SettingFile SettingStu
var SettingFileOnce sync.Once

func GetSettingFile() *SettingStu {
	SettingFileOnce.Do(SettingFile.Init)
	return &SettingFile
}

type SettingStu struct {
	XMLName  xml.Name      `xml:"setting"`
	Id       string        `xml:"settingid"`
	Cases    []*CaseStu    `xml:"case"`
	Aggre    AggregateIntf `xml:"-"`
	FilePath string        `xml:"-"`
}

func (r *SettingStu) Init() {
	r.FilePath = "./setting/setting.xml"
	ret := r.ReadFile()
	if ret != a1.Success {
		a1.PrnLog.Error("readfile error")
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
		ret = a1.ErrSettingReadFile
		return
	}

	err = xml.Unmarshal(content, r)
	if err != nil {
		a1.PrnLog.Error("ReadFile error", err)
		ret = a1.ErrSettingXmlFormat
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
