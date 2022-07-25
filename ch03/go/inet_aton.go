package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"
	"unsafe"
)

func IsLittleEndian() bool {
	var x int32 = 1 // 占 4byte 转换成16进制 0x00 00 00 01
	// 大端（16进制）：00 00 00 01
	// 小端（16进制）：01 00 00 00
	pointer := unsafe.Pointer(&x)
	pb := (*byte)(pointer)
	if *pb != 1 {
		return false
	}
	return true
}

// little endian to big endian
func ltob32(x uint32) uint32 {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, x)
	x = binary.LittleEndian.Uint32(bytes)
	return x
}

// host to network
func hton32(x uint32) uint32 {
	if IsLittleEndian() {
		x = ltob32(x)
	}
	return x
}

// IPstring to integer 把IP字符串转为数值
func aton32(ip string) (uint32, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, errors.New("invalid ipv4 format")
	}

	i := uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24

	i = hton32(i) // 网络字节序使用大端序

	return i, nil
}

func main() {
	addr := "127.232.124.79"

	net_addr, err := aton32(addr)
	if err != nil {
		checkError(err)
	}

	fmt.Printf("Network ordered integer addr: 0x%x\n", net_addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s ", err.Error())
		os.Exit(1)
	}
}
