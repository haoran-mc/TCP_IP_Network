package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const BUF_SIZE int = 30

// 获取空闲端口
func getAvailablePort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}

	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip port ", os.Args[0])
		os.Exit(1)
	}

	randPort, err := getAvailablePort()
	if err != nil {
		checkError(err)
	}

	listener, err := net.ListenPacket("udp4", fmt.Sprintf("%d", randPort))
	if err != nil {
		checkError(err)
	}

	message := make([]byte, BUF_SIZE)

	for {
		fmt.Print("Insert message(q to quit): ")
		str_len, _ := fmt.Scanln(&message)
		if strings.ToUpper(string(message[:str_len])) == "Q" {
			break
		}

		raddr, _ := net.ResolveUDPAddr("udp", os.Args[1]+":"+os.Args[2])

		listener.WriteTo(message, raddr)
		_, _, _ = listener.ReadFrom(message)

		fmt.Printf("Message from server: %s\n", message)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
