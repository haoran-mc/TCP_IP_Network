package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("udp", "224.1.1.2:9190")
	if err != nil {
		checkError(err)
	}

	for {
		conn.Write([]byte("hello, world!"))
		time.Sleep(1 * time.Second)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
