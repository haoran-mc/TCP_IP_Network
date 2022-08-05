package main

import (
	"fmt"
	"net"
	"os"
)

const BUF_SIZE int = 8192

func main() {
	// group addr
	gaddr, _ := net.ResolveUDPAddr("udp", "224.1.1.2:9190")
	listener, err := net.ListenMulticastUDP("udp", nil, gaddr)
	if err != nil {
		checkError(err)
	}

	// listener.SetReadBuffer(maxDatagramSize)

	message := make([]byte, BUF_SIZE)

	for {
		n, src, _ := listener.ReadFromUDP(message)
		fmt.Println(src, ": ", string(message[:n]))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
