package mux

import (
	"fmt"
	"net/http"
	"sync"

	. "base"

	Tool "tool"
)

var MuxSys MuxSysStu
var MuxSysOnce sync.Once

func GetMuxSys() http.Handler {
	MuxSysOnce.Do(MuxSys.Init)
	return &MuxSys
}

type MuxSysStu struct {
	http.ServeMux
}

func (r *MuxSysStu) Init() {
	r.HandleFunc("/", SysHandle)
}

func SysHandle(w http.ResponseWriter, req *http.Request) {
	var ret int
	var EntryResp *EntryStu
	defer func() {
		Tool.SetHeader(w)
		if ret == 0 {
			if EntryResp != nil {
				w.Write([]byte(EntryResp.GetBody()))
			}
		}
	}()
	Factory := GetSingleFactory()
	Problem := fmt.Sprintf("%v", req.URL.Path)
	fmt.Println("problem ---> ", Problem)
	Entry := Factory.Create(Problem)
	ret = Entry.Init(req)
	if ret != 0 {
		return
	}
	EntryResp, ret = Entry.Execute()
	if ret != 0 {

	}
	return
}
