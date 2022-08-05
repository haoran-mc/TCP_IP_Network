package main

import (
	"fmt"
	"net"
	"os"
)

const BUF_SIZE int = 1024

func main() {
	listener, err := net.ListenPacket("udp4", ":8888")
	if err != nil {
		checkError(err)
	}

	message := make([]byte, BUF_SIZE)

	for {
		n, src, err := listener.ReadFrom(message)
		if err != nil {
			checkError(err)
		}
		fmt.Println(src, ": ", string(message[:n]))
	}

	// listener.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
