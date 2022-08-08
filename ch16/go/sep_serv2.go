package main

/*
#include <sys/socket.h>
*/
import "C"
import (
	"fmt"
	"net"
	"os"
	"syscall"
)

const BUF_SIZE int = 1024

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s :port ", os.Args[0])
		os.Exit(1)
	}

	listener, _ := net.Listen("tcp4", os.Args[1])

	conn, _ := listener.Accept()

	tcpConn, _ := conn.(*net.TCPConn)
	fp, _ := tcpConn.File()
	fd := fp.Fd()

	newfd, _ := syscall.Dup(int(fd)) // 复制文件描述符

	readfp := os.NewFile(fd, "")
	writefp := os.NewFile(uintptr(newfd), "")

	writefp.WriteString("FROM SERVER: Hi~ client?\n")
	writefp.WriteString("I love all of the world \n")
	writefp.WriteString("You are awesome! \n")

	C.shutdown(C.int(writefp.Fd()), C.SHUT_WR) // 调用 C 函数：shutdown
	writefp.Close()                            // 关闭写流

	buf := make([]byte, BUF_SIZE)
	n, _ := readfp.Read(buf)

	fmt.Println(string(buf[:n]))
	readfp.Close() // 关闭读流
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
