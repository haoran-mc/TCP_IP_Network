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

// 将 int32 类型的变量转换为 []byte 类型
// 放在 bys 的 idx 位置，占长度为 length
func BytesAddInt32(bys []byte, n int32, idx int) []byte {
	s := Int32ToBytes(n)
	for i := 0; i < 4; i++ {
		bys[idx+i] = s[i]
	}
	return bys
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip port ", os.Args[0])
		os.Exit(1)
	}

	conn, _ := net.Dial("tcp4", os.Args[1]+":"+os.Args[2])

	var opnd_cnt int32
	opmsg := make([]byte, BUF_SIZE)

	fmt.Print("Operand count: ")
	fmt.Scanf("%d", &opnd_cnt)
	// 在 opmsg 的第一个位置放待算数字的个数，占一个字节
	opmsg[0] = Int32ToBytes(opnd_cnt)[3]

	for i := int32(0); i < opnd_cnt; i++ {
		var opnum int32
		fmt.Printf("Operand %d: ", i+1)
		fmt.Scanf("%d", &opnum)
		// 在 []byte 中添加一个 int32 类型的待算数，占 4 个字节
		opmsg = BytesAddInt32(opmsg, opnum, int(i)*OPSZ+1)
	}

	fmt.Print("Operator: ")
	fmt.Scanf("%c", &opmsg[int(opnd_cnt)*OPSZ+1]) // 最后一个字节放运算符

	// fmt.Println(opmsg[:int(opnd_cnt)*OPSZ+2])

	conn.Write(opmsg[:int(opnd_cnt)*OPSZ+2])

	result := make([]byte, 4)
	conn.Read(result)

	fmt.Print("Result: ")

	fmt.Printf("Operation result: %d \n", BytesToInt32(result))
	conn.Close()
}
