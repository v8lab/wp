package base

import (
	"net"
)

type EntryIntf interface {
	Init(*net.UDPAddr, *net.UDPAddr, *net.UDPConn)
	Execute() int
}
