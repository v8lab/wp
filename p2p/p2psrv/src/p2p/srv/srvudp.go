package srv

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	mylib "mylib"
	base "mylib/udpentry/base"

	stun "github.com/ccding/go-stun/stun"
)

var SingleSrvUdp SrvUdpStu
var SingleSrvUdpOnce sync.Once

func GetSingleSrvUdp() *SrvUdpStu {
	SingleSrvUdpOnce.Do(SingleSrvUdp.Init)
	return &SingleSrvUdp
}

type SrvUdpStu struct {
	conn *net.UDPConn
	cli  *stun.Client
}

func (r *SrvUdpStu) Init() {
	addr, err := net.ResolveUDPAddr("udp", ":55500")
	if err != nil {
		mylib.PrnLog.Error("Can't resolve address: ", err)
		os.Exit(1)
	}
	r.conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		mylib.PrnLog.Error("net.ListenUDP")
		return
	}
	t := time.Now()
	r.conn.SetDeadline(t.Add(time.Duration(30 * time.Second)))
	r.cli = stun.NewClientWithConnection(r.conn)
	return
}
func (r *SrvUdpStu) Dis() {
	a, b, c := r.cli.Discover()
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
}
func (r *SrvUdpStu) GetConn() (conn *net.UDPConn) {
	return r.conn
}

func (r *SrvUdpStu) StartRead() {
	for {
		data := make([]byte, 4096)
		n, addr, err := r.conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println("err", err)
			continue
		} else {
			fmt.Println("addr", addr)
			fmt.Println("data", data[:n])
			r.WriteAddr([]byte("i am server"), addr)
			go base.EntryFacade(data[:n], addr)
		}
	}
}
func (r *SrvUdpStu) ConnTicker() {
	ticker := time.NewTicker(time.Millisecond * 50)
	for range ticker.C {
		t := time.Now()
		r.conn.SetDeadline(t.Add(time.Duration(5 * time.Second)))
	}
}
func (r *SrvUdpStu) WriteAddr(Data []byte, Addr net.Addr) {
	r.conn.WriteTo(Data, Addr)
	return
}
func (r *SrvUdpStu) WriteAddrStr(Data []byte, raddrstr string) {
	raddr, err := net.ResolveUDPAddr("udp", raddrstr)
	if err != nil {
		fmt.Println("net.ResolveUDPAddr fail.", err)
		return
	}
	r.WriteAddr(Data, raddr)
	return
}
