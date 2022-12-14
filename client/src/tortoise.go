package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
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
	// Create Dail To Login-创建Tcp连接并登录
	ip, err := stringToIP(host)
	if err != nil {
		return
	}
	raddr := net.TCPAddr{IP: ip, Port: port}
	tcpConn, err := net.DialTCP("tcp4", nil, &raddr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	_, err = tcpConn.Write(stringToBytes("hello"))
	if err != nil {
		return
	}
	// Cycling and Receiving Commands-循环接收指令
}

func stringToIP(host string) (net.IP, error) {
	section := strings.Split(host, ".")
	if len(section) != 4 {
		return nil, errors.New("host format error:" + host)
	}
	a, err := strconv.Atoi(section[0])
	if err != nil {
		return nil, err
	}
	b, err := strconv.Atoi(section[1])
	if err != nil {
		return nil, err
	}
	c, err := strconv.Atoi(section[2])
	if err != nil {
		return nil, err
	}
	d, err := strconv.Atoi(section[3])
	if err != nil {
		return nil, err
	}
	return net.IPv4(byte(a), byte(b), byte(c), byte(d)), nil
}

func stringToBytes(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}
