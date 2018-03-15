package mux

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	Tool "tool"
)

var SingleMuxCli MuxCliStu
var SingleMuxCliOnce sync.Once

func GetMuxCli() http.Handler {
	SingleMuxCliOnce.Do(SingleMuxCli.Init)
	return &SingleMuxCli
}

type MuxCliStu struct {
	http.ServeMux
}

func (r *MuxCliStu) Init() {
	r.HandleFunc("/", SysHandle)
}

func SysHandle(w http.ResponseWriter, req *http.Request) {

	var (
		ret int
	)

	defer func() {
		Tool.SetHeader(w)

		w.Header().Add("errcode", strconv.Itoa(ret))

	}()

	//	PrnLog.Debug()
	fmt.Println("\n\n\tclient response : \n\n")

	return
}
