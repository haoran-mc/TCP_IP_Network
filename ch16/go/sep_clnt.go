package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const BUF_SIZE int = 1024

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip port ", os.Args[0])
		os.Exit(1)
	}

	conn, _ := net.Dial("tcp4", os.Args[1]+":"+os.Args[2])

	tcpConn := conn.(*net.TCPConn)
	fp, _ := tcpConn.File()
	fd := fp.Fd()

	readfp := os.NewFile(fd, "")
	writefp := os.NewFile(fd, "")

	buf := make([]byte, BUF_SIZE)

	for {
		n, err := readfp.Read(buf)

		if n == 0 || err == io.EOF {
			break
		}

		fmt.Println(string(buf[:n]))
	}

	writefp.WriteString("FROM CLIENT: Thank you \n")
	writefp.Close()
	readfp.Close()
}
