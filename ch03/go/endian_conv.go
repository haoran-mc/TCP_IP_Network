package main

import (
	"encoding/binary"
	"fmt"
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
func ltob16(x uint16) uint16 {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, x) // bytes 中存放的是大端序
	// 不经过转换，在小端序主机上输出从网络中获取到的字节序，得到的值：
	x = binary.LittleEndian.Uint16(bytes)
	return x
}

// host to network
func hton16(x uint16) uint16 {
	if IsLittleEndian() {
		x = ltob16(x)
	}
	return x
}

func ltob32(x uint32) uint32 {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, x)
	x = binary.LittleEndian.Uint32(bytes)
	return x
}

func hton32(x uint32) uint32 {
	if IsLittleEndian() {
		x = ltob32(x)
	}
	return x
}

func main() {
	// 0000 0000  0000 0000  0001 0010  0011 0100
	var host_port uint16 = 0x1234
	var net_port uint16
	// 0001 0010  0011 0100  0101 0110  0111 1000
	var host_addr uint32 = 0x12345678
	var net_addr uint32

	net_port = hton16(host_port)
	net_addr = hton32(host_addr)

	fmt.Printf("Host ordered port: 0x%x\n", host_port)
	fmt.Printf("Network ordered port: 0x%x\n", net_port)
	fmt.Printf("Host ordered port: 0x%x\n", host_addr)
	fmt.Printf("Network ordered port: 0x%x\n", net_addr)
}
