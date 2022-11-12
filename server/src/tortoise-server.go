package main

import (
	"log"
	"net"
)

func init() {

}

func main() {
	// define listener
	listener, err := net.Listen("tcp", "127.0.0.1:56701")
	if err != nil {
		log.Fatalln("tortoise server start error: ", err)
		return
	}
	//defer listener.Close()
	log.Println("tortoise server is starting...")

	for {
		c, err := listener.Accept()
		if err != nil {
			log.Fatalln("accept error:", err)
			//break
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	for {
		var buf = make([]byte, 10)
		read, err := c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		log.Printf("read %d bytes, content is %s\n", read, string(buf[:read]))
	}
}
