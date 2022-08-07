package main

import (
	"os"
	"syscall"
)

func main() {
	fp, _ := os.OpenFile("data2.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)

	fd := fp.Fd() // 获取 os.File 结构体指针表示的文件描述符

	syscall.Write(int(fd), []byte("TCP/IP SOCKET PROGRAMMING\n"))

	fp.Close()
}
