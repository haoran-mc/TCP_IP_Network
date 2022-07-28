package main

import (
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

	listener, err := net.ListenPacket("udp4", os.Args[1])
	if err != nil {
		checkError(err)
	}

	message := make([]byte, BUF_SIZE)

	for {
		str_len, raddr, err := listener.ReadFrom(message)
		if err != nil {
			checkError(err)
		}
		listener.WriteTo(message[:str_len], raddr)
	}

	// listener.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
