package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

const (
	BUF_SIZE = 100
	MAX_CLNT = 256
)

var clnt_cnt = 0
var clnt_socks = make([]net.Conn, MAX_CLNT)
var mutex sync.Mutex // 创建互斥锁

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s :port ", os.Args[0])
		os.Exit(1)
	}

	listener, err := net.Listen("tcp4", os.Args[1])
	if err != nil {
		checkError(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			checkError(err)
		}

		mutex.Lock() // 上锁
		clnt_socks[clnt_cnt] = conn
		clnt_cnt++
		mutex.Unlock() // 解锁

		go handle_clnt(conn)
		fmt.Printf("Connected client IP: %s\n", conn.RemoteAddr().String())
	}
	// listener.Close()
}

func handle_clnt(conn net.Conn) {
	var msg = make([]byte, BUF_SIZE)

	for {
		str_len, err := conn.Read(msg)
		if err == io.EOF { // 断开连接
			break
		}

		send_msg(conn, msg, str_len)
	}

	mutex.Lock()
	for i := 0; i < clnt_cnt; i++ {
		if clnt_socks[i] == conn {
			for {
				if i >= clnt_cnt-1 {
					break
				}
				i++
				clnt_socks[i] = clnt_socks[i+1]
			}
			break
		}
	}
	clnt_cnt--
	mutex.Unlock()

	conn.Close()
}

func send_msg(conn net.Conn, msg []byte, str_len int) {
	mutex.Lock()
	for i := 0; i < clnt_cnt; i++ {
		clnt_socks[i].Write(msg[:str_len])
	}
	mutex.Unlock()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
