package handler

import (
	"bufio"
	"log"
	"net"
	"strconv"
	//"server/utils"
	//"strconv"
)

func HandleConnection(tcpConn *net.TCPConn) {
	remoteAddr := tcpConn.RemoteAddr()
	defer tcpConn.Close()
	writer := bufio.NewWriter(tcpConn)

	for {
		data := make([]byte, 2048)
		//bufio.NewReader(tcpConn)
		n, err := tcpConn.Read(data)
		if err != nil {
			log.Printf("remoteAddr=%s, error=%s\n", remoteAddr, err.Error())
			tcpConn.Write([]byte("Error"))
			break
		}
		log.Printf("remoteAddr=%s, data=%s\n", remoteAddr, string(data))

		_, err = writer.WriteString("Receive " + strconv.Itoa(n) + " bytes")
		if err != nil {
			return
		}
		err = writer.Flush()
		if err != nil {
			return
		}
		//output := "Receive " + strconv.Itoa(n) + " bytes"
		//_, err = tcpConn.Write(data)
		//if err != nil {
		//	log.Printf("remoteAddr=%s, error=%s\n", remoteAddr, err.Error())
		//	continue
		//}
		log.Println("Write To Client")
	}
}
