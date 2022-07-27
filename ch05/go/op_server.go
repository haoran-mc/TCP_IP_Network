package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

const (
	BUF_SIZE int = 1024
	OPSZ     int = 4
)

// 将 int32 型转化为 []byte 类型
func Int32ToBytes(n int32) []byte {
	data := int32(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

// 将 []byte 类型的变量转换为 int32 类型
func BytesToInt32(bys []byte) int32 {
	bytebuff := bytes.NewBuffer(bys)
	var data int32
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int32(data)
}

func calculate(opnum int32, opnds []int32, op byte) int32 {
	result := opnds[0]
	switch op {
	case '+':
		for i := int32(1); i < opnum; i++ {
			result += opnds[i]
		}
		break
	case '-':
		for i := int32(1); i < opnum; i++ {
			result -= opnds[i]
		}
		break
	case '*':
		for i := int32(1); i < opnum; i++ {
			result *= opnds[i]
		}
		break
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s :port ", os.Args[0])
		os.Exit(1)
	}

	listener, err := net.Listen("tcp4", os.Args[1])
	if err != nil {
		checkError(err)
	}

	for i := 1; i < 5; i++ {
		conn, err := listener.Accept()
		if err != nil {
			checkError(err)
		}

		opnd_cnt_byte := make([]byte, 1)
		conn.Read(opnd_cnt_byte) // 待算数字个数，只接收一个字节，也就是第一个字节
		opnd_cnt := int32(opnd_cnt_byte[0])

		opinfo := make([]byte, BUF_SIZE)
		recv_len, recv_cnt := 0, 0

		// 每个整数占4个字节 + 操作符1个字节 == 接收到的字节数时，停止读取
		for int(opnd_cnt)*OPSZ+1 > recv_len {
			recv_cnt, _ = conn.Read(opinfo)
			recv_len += recv_cnt
		}

		// calculate 的参数是 []int32，所以需要对 opinfo 转换为 []int32类型
		opnums := make([]int32, opnd_cnt)

		for i := 0; i < int(opnd_cnt); i++ {
			opnums[i] = BytesToInt32(opinfo[i*OPSZ : (i+1)*OPSZ])
		}

		result := calculate(opnd_cnt, opnums, opinfo[recv_len-1])
		conn.Write(Int32ToBytes(result))
		conn.Close()
	}
	listener.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
