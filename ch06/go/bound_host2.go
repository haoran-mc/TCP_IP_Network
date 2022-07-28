package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	msg1 := []byte("Hi!")
	msg2 := []byte("I'm another UDP host!")
	msg3 := []byte("Nice to meet you")

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip port ", os.Args[0])
		os.Exit(1)
	}

	conn, err := net.Dial("udp4", os.Args[1]+":"+os.Args[2])
	if err != nil {
		checkError(err)
	}

	conn.Write(msg1)
	conn.Write(msg2)
	conn.Write(msg3)

	conn.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
