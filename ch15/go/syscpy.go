package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

const BUF_SIZE int = 3

func main() {
	start := time.Now()

	fd1, _ := syscall.Open("data.txt", os.O_RDONLY, 0755)
	fd2, _ := syscall.Open("syscpy.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)

	buf := make([]byte, BUF_SIZE)

	for {
		n, _ := syscall.Read(fd1, buf)
		if n == 0 {
			break
		}

		syscall.Write(fd2, buf[:n])
	}

	syscall.Close(fd1)
	syscall.Close(fd2)

	elapsed := time.Since(start)
	fmt.Println("The procedure takes ", elapsed, "s")
}
