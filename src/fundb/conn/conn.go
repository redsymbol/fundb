package conn

import (
	"strings"
	"net"
)

type Conn struct {
	Port int
	handle net.Conn
}

func (conn Conn) Send(msg string) string {
	
	return strings.Join([]string{
		"SENDING:",
		msg,
	}, " ")
}

func New(port int) *Conn {
	conn = Conn{Port:port}
	
}