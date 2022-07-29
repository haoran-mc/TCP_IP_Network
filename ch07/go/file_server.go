package main

/*
#include <sys/socket.h>
*/
import "C"

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const BUF_SIZE int = 30

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s :port ", os.Args[0])
		os.Exit(1)
	}

	listener, err := net.Listen("tcp4", os.Args[1])
	if err != nil {
		checkError(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		checkError(err)
	}

	fp, _ := os.Open("file_server.go")

	reader := bufio.NewReader(fp)
	buf := make([]byte, BUF_SIZE)

	for {
		read_cnt, _ := reader.Read(buf) // 读文件 file_server.go 内容
		if read_cnt < BUF_SIZE {
			conn.Write(buf[:read_cnt])
			break
		}
		conn.Write(buf) // 文件内容发送到客户端
	}

	tcpConn, _ := conn.(*net.TCPConn)
	tcpfp, _ := tcpConn.File()
	defer tcpfp.Close()
	tcpfd := tcpfp.Fd() // 获取到被 conn 封装起来的文件描述符

	C.shutdown(C.int(tcpfd), C.SHUT_WR) // 调用 C 函数：shutdown

	n, _ := conn.Read(buf) // 半关闭后仍然可以读 socket
	fmt.Printf("Message from client: %s \n", buf[:n])

	/*
		// 因为半关闭，所以下面代码不能将信息传递到客户端
		conn.Write([]byte("END"))
	*/

	fp.Close()
	conn.Close()
	listener.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
