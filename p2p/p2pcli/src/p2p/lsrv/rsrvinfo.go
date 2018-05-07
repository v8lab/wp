package lsrv

import (
	"fmt"
	"net"

	"sync"
)

var SingleRSrvInfo RSrvInfoStu
var SingleRSrvInfoOnce sync.Once

func GetSingleRSrvInfo() *RSrvInfoStu {
	SingleRSrvInfoOnce.Do(SingleRSrvInfo.Init)
	return &SingleRSrvInfo
}

type RSrvInfoStu struct {
	Id   int
	Addr *net.UDPAddr
}

func (r *RSrvInfoStu) Init() {
	raddrstr := "180.76.119.248:55500"
	raddr, err := net.ResolveUDPAddr("udp", raddrstr)
	if err != nil {
		fmt.Println("net.ResolveUDPAddr fail.", err)
		return
	}
	r.Addr = raddr
	return
}

func (r *RSrvInfoStu) GetAddr() (Addr *net.UDPAddr) {
	return r.Addr
}
