package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
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

// big endian to little endian
func btol32(x uint32) uint32 {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, x)
	x = binary.BigEndian.Uint32(bytes)
	return x
}

// network to host
func ntoh32(x uint32) uint32 {
	if IsLittleEndian() {
		x = btol32(x)
	}
	return x
}

// integer to IPstring
func ntoa32(net_addr uint32) (string, error) {
	if net_addr > math.MaxUint32 {
		return "", errors.New("beyond the scope of ipv4")
	}

	i := ntoh32(net_addr) // 转换为本地字节序的 ip 地址

	// 1.2.3.4
	// 0000 00001  0000 0010  0000 0011  0000 01000

	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)

	return ip.String(), nil
}

func main() {
	net_addr := hton32(0x1020304) // 网络中传来的大端序表示的 ip 地址

	ip, err := ntoa32(net_addr)
	if err != nil {
		checkError(err)
	}

	fmt.Printf("Dotted-Decimal notation1: %s\n", ip)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s ", err.Error())
		os.Exit(1)
	}
}
