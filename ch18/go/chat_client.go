package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
)

const (
	BUF_SIZE  = 100
	NAME_SIZE = 20
)

var name = make([]byte, NAME_SIZE)
var msg = make([]byte, BUF_SIZE)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip port name ", os.Args[0])
		os.Exit(1)
	}

	name = []byte(fmt.Sprintf("[%s]", os.Args[3]))

	conn, err := net.Dial("tcp4", os.Args[1]+":"+os.Args[2])
	if err != nil {
		checkError(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go send_msg(conn, &wg) // 创建发送消息协程
	go recv_msg(conn, &wg) // 创建接收消息协程

	wg.Wait()
}

func send_msg(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	name_msg := make([]byte, NAME_SIZE+BUF_SIZE)

	for {
		fmt.Scanf("%s", &msg)
		if strings.ToUpper(string(msg)) == "Q" {
			conn.Close()
			break
		}
		name_msg = []byte(fmt.Sprintf("%s %s", name, msg))
		conn.Write(name_msg)
	}
}

func recv_msg(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	name_msg := make([]byte, NAME_SIZE+BUF_SIZE)

	for {
		str_len, err := conn.Read(name_msg)
		if err == io.EOF {
			break
		}

		fmt.Println(string(name_msg[:str_len]))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
