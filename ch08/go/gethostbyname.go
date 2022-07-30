package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s Domain name ", os.Args[0])
		os.Exit(1)
	}

	ips, err := net.LookupHost(os.Args[1])
	if err != nil {
		checkError(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
