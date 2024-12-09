package tcpserver

import "net"

type Client struct {
	Conn         net.Conn
	IncomingChan chan []byte
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		Conn:         conn,
		IncomingChan: make(chan []byte, 100),
	}
}
