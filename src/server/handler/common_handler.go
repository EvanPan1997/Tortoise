package handler

import (
	"fmt"
	"net"
)

func HandleConnection(tcpConn net.TCPConn) {
	// 1.Verify Login
	// 2.Process Request
	// 3.Exit
}

func Handle(conn net.Conn) {
	//log.Println(conn.RemoteAddr())
	//log.Println(conn.LocalAddr())
	defer conn.Close()

	for {
		data := make([]byte, 2048)
		_, err := conn.Read(data)
		if err != nil {
			return
		}
		fmt.Println(string(data))
	}

}
