package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip port ", os.Args[0])
		os.Exit(1)
	}

	idx, read_len, str_len := 0, 0, 0

	conn, err := net.Dial("tcp4", os.Args[1]+":"+os.Args[2])
	checkError(err)

	// 每次只接收一个字符
	buf := make([]byte, 1)
	message := make([]byte, 1024)
	for {
		read_len, err = conn.Read(buf)
		if err == io.EOF { // 这里的用法查看 conn.Read() 码源
			break
		}
		str_len += read_len
		message[idx] = buf[0]
		idx++
	}
	fmt.Printf("Message from server: %s \n", string(message))
	fmt.Printf("Function read call count: %d \n", str_len)
	conn.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s ", err.Error())
		os.Exit(1)
	}
}
