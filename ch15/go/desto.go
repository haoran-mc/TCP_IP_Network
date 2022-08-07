package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
)

func main() {
	fd, _ := syscall.Open("data1.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)

	fp := os.NewFile(uintptr(fd), "") // 将文件描述符转换为 os.File 结构体指针

	writer := bufio.NewWriter(fp)
	writer.WriteString(fmt.Sprintf("NetWork Go programming %v", fp))
	writer.Flush()

	fp.Close()

	fmt.Println(fp)
}
