package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip ", os.Args[0])
		os.Exit(1)
	}

	addrs, err := net.LookupAddr(os.Args[1])
	if err != nil {
		checkError(err)
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
