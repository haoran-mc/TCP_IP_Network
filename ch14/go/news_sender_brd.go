package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	/*
		// 这里设置发送者的IP地址，自己查看一下自己的IP自行设定
		laddr := net.UDPAddr{
			IP:   net.IPv4(192, 168, 137, 224),
			Port: 3000,
		}
		// 这里设置接收者的IP地址为广播地址
		raddr := net.UDPAddr{
			IP:   net.IPv4(255, 255, 255, 255),
			Port: 3000,
		}
		conn, err := net.DialUDP("udp", &laddr, &raddr)
	*/

	conn, err := net.Dial("udp4", "255.255.255.255:8888")
	if err != nil {
		checkError(err)
	}

	for {
		conn.Write([]byte("hello world"))
		time.Sleep(1 * time.Second)
	}

	// conn.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
