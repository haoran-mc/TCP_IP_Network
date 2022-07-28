package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const BUF_SIZE int = 30

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip port ", os.Args[0])
		os.Exit(1)
	}

	conn, err := net.Dial("udp4", os.Args[1]+":"+os.Args[2])
	if err != nil {
		checkError(err)
	}

	message := make([]byte, BUF_SIZE)

	for {
		fmt.Print("Insert message(q to quit): ")
		str_len, _ := fmt.Scanln(&message)
		if strings.ToUpper(string(message[:str_len])) == "Q" {
			break
		}

		conn.Write(message[:str_len])
		conn.Read(message)

		fmt.Printf("Message from server: %s\n", message)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
