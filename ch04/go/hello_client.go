package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}

	// 监听端口由命令行参数得到
	service := os.Args[1]

	conn, err := net.Dial("tcp4", service)
	checkError(err)

	buf := make([]byte, 1024)
	// 读取来自服务端的数据，返回数据长度
	n, err := conn.Read(buf)
	checkError(err)

	fmt.Println(string(buf[:n]))
	conn.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
