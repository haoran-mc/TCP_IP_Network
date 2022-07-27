package main

import (
	"fmt"
	"net"
	"os"
)

const BUF_SIZE int = 1024

func main() {
	message := make([]byte, BUF_SIZE)
	var str_len, recv_len, recv_cnt int

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip port ", os.Args[0])
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

		str_len, _ = conn.Write(message)

		for {
			if recv_len >= str_len {
				break
			}
			recv_cnt, err = conn.Read(message[recv_len:])
			if err != nil {
				checkError(err)
			}
			recv_len += recv_cnt
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
