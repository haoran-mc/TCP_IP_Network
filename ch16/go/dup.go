package main

import (
	"fmt"
	"syscall"
)

func main() {
	str1 := "Hi! \n"
	str2 := "It's nice day! \n"

	/*
		文件描述符：
		1 - 标准输入
		2 - 标准输出
		3 - 标准错误
	*/
	cfd1, _ := syscall.Dup(1)
	var cfd2 int = 7
	syscall.Dup2(cfd1, cfd2) // 复制文件描述符 cfd1 得到 cfd2

	fmt.Printf("fd1 = %d, fd2 = %d \n", cfd1, cfd2)

	syscall.Write(cfd1, []byte(str1)) // 向 cfd1 和 cfd2 写入，等同于向标准输出写入
	syscall.Write(cfd2, []byte(str2))

	syscall.Close(cfd1)
	syscall.Close(cfd2)

	syscall.Write(1, []byte(str1))
	syscall.Close(1)
	syscall.Write(1, []byte(str2))
}
