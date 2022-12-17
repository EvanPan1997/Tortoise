package main

import (
	"net"
	"server/handler"
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
		//conn, err := tcpListener.Accept()
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			return
		}
		//go handler.Handle(conn)
		go handler.HandleConnection(conn)

		//tcpConn, err := tcpListener.AcceptTCP()
		//if err != nil {
		//	return
		//}
		//tcpConn.SetKeepAlive(true)
		//tcpConn.SetKeepAlivePeriod(time.Minute * 5)
		//tcpConn.SetLinger(10)
		//go handler.HandleConnection(tcpConn)
	}
}
