package main

import (
	"fmt"
	"net"
	"os"
)

const BUF_SIZE int = 1024

func main() {
	message := make([]byte, BUF_SIZE)
	var str_len int

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <ip> <port> ", os.Args[0])
		os.Exit(1)
	}

	conn, err := net.Dial("tcp4", os.Args[1]+":"+os.Args[2])
	if err != nil {
		checkError(err)
	}

	for {
		fmt.Print("Input message(Q to quit): ")
		fmt.Scanf("%s", &message)

		if string(message) == "q" || string(message) == "Q" {
			break
		}

		str_len, err = conn.Write(message)
		if err != nil {
			checkError(err)
		}

		message[0] = 0 // 确保 message 里面的数据是从服务方发来的

		str_len, err = conn.Read(message)
		if err != nil {
			checkError(err)
		}
		fmt.Printf("Message from server: %s\n", message[:str_len])
	}

	conn.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
