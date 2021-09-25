package simpleTcpMessage

import "net"

type ClientConn struct {
	conn net.Conn
}

func NewClientConn(conn net.Conn) *ClientConn {
	return &ClientConn{
		conn: conn,
	}
}
