package mux

import (
	"net/http"
	"strconv"
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

	defer func() {
		Tool.SetHeader(w)

		w.Header().Add("errcode", strconv.Itoa(ret))
		w.Header().Add("errmsg", "")

		if ret == 0 {

		} else {
			w.Write([]byte(""))
		}
	}()

	Entry := NewEntryKind("")
	ret = Entry.Init(req)
	if ret != 0 {
		return
	}
	ret = Entry.Execute()
	if ret != 0 {

	}
	return
}
