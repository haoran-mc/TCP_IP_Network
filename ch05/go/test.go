package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// 将 int 型转化为 []byte 类型
func IntToBytes(n int) []byte {
	data := int32(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func main() {
	fmt.Print(IntToBytes(1))
}
