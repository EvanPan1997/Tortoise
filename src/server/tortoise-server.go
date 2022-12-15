package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// default enter, analysis path and choose an available handler
	laddr := net.TCPAddr{
		IP:   nil,
		Port: 9890,
	}
	tcpListener, err := net.ListenTCP("tcp4", &laddr)
	if err != nil {
		return
	}
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			return
		}
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(time.Minute * 5)
		tcpConn.SetLinger(10)
		go handlerConnection(tcpConn)
	}
}

func handlerConnection(tcpConn *net.TCPConn) {
	fmt.Println("success")
	// 1.Verify Login
	// 2.Process Request
	// 3.Exit
}
