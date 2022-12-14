package main

import (
	"net"
)

func main() {
	// default enter, analysis path and choose an available handler
	laddr := &net.TCPAddr{
		IP:   nil,
		Port: 9890,
	}
	tcpListener, err := net.ListenTCP("tcp4", laddr)
	if err != nil {
		return
	}
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			return
		}
		go handlerConnection(tcpConn)
	}
}

func handlerConnection(listener *net.TCPConn) {

}
