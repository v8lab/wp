package setting

import (
	"encoding/xml"

	. "httpsend/intfscan"
)

func NewCaseStu() *CaseStu {
	var Case CaseStu
	Case.Init()
	return &Case
}

type CaseStu struct {
	XMLName xml.Name      `xml:"case"`
	Id      string        `xml:"caseid"`
	Units   []*UnitStu    `xml:"unit"`
	Aggre   AggregateIntf `xml:"-"`
}

func (r *CaseStu) Init() {
	r.Aggre = NewAggregateStu()
	for _, v := range r.Units {
		vv := v
		vv.Init()
		r.Aggre.Add(vv.GetId(), vv)
	}
}

func (r *CaseStu) GetData() interface{} {
	return r
}

func (r *CaseStu) Update() (ret int) {
	return
}

func (r *CaseStu) GetId() string {
	return r.Id
}
func (r *CaseStu) Find(Id string) (o *UnitStu, exist bool) {
	v, e := r.Aggre.Find(Id)
	if e {
		if vv, ok := v.(*UnitStu); ok {
			o = vv
			exist = true
		}
	}
	return
}

func (r *CaseStu) Iterator() IterIntf {
	return r.Aggre.Iterator()
}
