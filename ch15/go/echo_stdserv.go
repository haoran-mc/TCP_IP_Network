package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

const BUF_SIZE int = 1024

func main() {
	message := make([]byte, BUF_SIZE)

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s :port ", os.Args[0])
		os.Exit(1)
	}

	listener, err := net.Listen("tcp4", os.Args[1])
	if err != nil {
		checkError(err)
	}

	// 调用 5 次 Accept() 方法，共为 5 个客户端提供服务
	for i := 0; i < 5; i++ {
		conn, err := listener.Accept()
		if err != nil {
			checkError(err)
		}

		fmt.Printf("Connect client %d \n", i+1)

		tcpConn, _ := conn.(*net.TCPConn)
		fp, _ := tcpConn.File() // 获取 tcp 链接对应的 os.File
		reader := bufio.NewReader(fp)
		writer := bufio.NewWriter(fp)

		for {
			str_len, err := reader.Read(message)
			if err == io.EOF {
				break
			}

			fmt.Print(conn.RemoteAddr().String(), "  ")
			fmt.Println(string(message[:str_len]))

			writer.Write(message[:str_len])
			writer.Flush()
		}
		conn.Close()
	}
	listener.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
