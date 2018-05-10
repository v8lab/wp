package setting

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"

	"sync"
)

var Setting SettingStu
var SettingOnce sync.Once

func GetSetting() *SettingStu {
	SettingOnce.Do(Setting.Init)
	return &Setting
}

type SettingStu struct {
	XMLName xml.Name `xml:"setting"`
	Port    string   `xml:"port"`

	sync.RWMutex `xml:"-"`
	FilePath     string `xml:"-"`
}

func (r *SettingStu) Init() {
	r.FilePath = "./setting.xml"
	ret := r.ReadFile()
	if ret != 0 {
		fmt.Println("readfile error")
		return
	}
	return
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

func (r *SettingStu) ToJson() string {
	jsonBody, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(jsonBody)
}
