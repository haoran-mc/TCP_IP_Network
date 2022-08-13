package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const BUF_SIZE int = 30

// 文件是否存在
func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip port ", os.Args[0])
		os.Exit(1)
	}

	filename := "receive.txt"
	var fp *os.File

	if checkFileIsExist(filename) { // 判断文件是否存在
		fp, _ = os.OpenFile(filename, os.O_APPEND, 0666)
	} else {
		fp, _ = os.Create(filename)
	}

	conn, err := net.Dial("tcp4", os.Args[1]+":"+os.Args[2])
	if err != nil {
		checkError(err)
	}

	buf := make([]byte, BUF_SIZE)

	for {
		read_cnt, err := conn.Read(buf) // 接收来自服务端的信息
		if err == io.EOF || read_cnt < BUF_SIZE {
			break
		}
		fp.Write(buf[:read_cnt]) // 写入文件
	}

	fmt.Println("Received file data")
	// 此时服务端半关闭了，仍然发送消息，观察服务端能否接收到数据
	conn.Write([]byte("Thank you"))

	/*
		// 因为半关闭，所以下面代码不能收到服务端的信息
		n, _ := conn.Read(buf)
		fmt.Println("Message after shutdown: ", string(buf[:n]))
	*/

	fp.Close()
	conn.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
