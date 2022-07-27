package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s :port ", os.Args[0])
		os.Exit(1)
	}
	// 监听端口由命令行参数得到
	service := os.Args[1]

	// 1. 监听请求 Listen
	listener, err := net.Listen("tcp4", service)
	checkError(err)

	// 2. 接收请求 Accept
	conn, err := listener.Accept()
	checkError(err)

	message := []byte("hello world")
	conn.Write(message)
	conn.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
