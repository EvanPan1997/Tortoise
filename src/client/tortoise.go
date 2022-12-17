package main

import (
	"client/utils"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Client Initialize-客户端初始化
	var (
		defaultUser string = "root"
		defaultHost string = "127.0.0.1"
		defaultPort int    = 9890

		user     string
		password string
		host     string
		port     int
	)
	// Receive Arguments From Command Line-接收命令行传参
	flag.StringVar(&user, "u", defaultUser, "")
	flag.StringVar(&host, "h", defaultHost, "")
	flag.IntVar(&port, "p", defaultPort, "")
	// Analysis Arguments-参数解析
	flag.Parse()

	// Receive Password To Login-接收密码用于登录
	n := 0
	for password == "" {
		if n <= 2 {
			fmt.Printf("Please enter the password: ")
			_, err := fmt.Scan(&password)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(-1)
			}
		} else {
			os.Exit(-1)
		}
		n++
	}
	// Check The Length of Password-校验密码长度
	if len(password) < 8 {
		fmt.Println("The length of password is too low")
		os.Exit(-1)
	}

	raddr := net.TCPAddr{IP: net.ParseIP(host), Port: port}
	tcpConn, err := net.DialTCP("tcp", nil, &raddr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	var input string
	// Cycling and Receiving Commands-循环接收指令
	for {
		fmt.Printf(">")
		_, inputErr := fmt.Scanf("%s", &input)
		if inputErr != nil {
			log.Fatalln(inputErr)
			//os.Exit(100)
		}

		input = strings.TrimSpace(input)
		// 排除空串
		if input != "" {
			transportToServer(tcpConn, input)
		}
	}

}

func transportToServer(tcpConn *net.TCPConn, input string) {
	// 发送请求
	_, err := tcpConn.Write(utils.StringToBytes(input))
	if err != nil {
		os.Exit(101)
	}
	// 接收返回
	output := make([]byte, 2048)
	//reader := bufio.NewReader(tcpConn)
	//n, err := reader.Read(output)
	//if err != nil {
	//	return
	//}
	n, err := tcpConn.Read(output)
	if err != nil {
		os.Exit(102)
	}
	fmt.Printf("Receive %d, data=%s\n", n, string(output))
}
