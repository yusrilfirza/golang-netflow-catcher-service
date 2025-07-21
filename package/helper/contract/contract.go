package contract

import (
	"net"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Log(level logrus.Level, message string, field interface{}, err error)
}

type Intfrastructure interface {
	ResolveUDPAddr(port string) (*net.UDPAddr, error)

	ListenUDP(address *net.UDPAddr) (*net.UDPConn, error)
}
