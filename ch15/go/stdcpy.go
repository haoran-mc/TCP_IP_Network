package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const BUF_SIZE int = 3

func main() {
	start := time.Now()

	fp1, _ := os.OpenFile("data.txt", os.O_RDONLY, 0755)
	fp2, _ := os.OpenFile("stdcpy.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)

	reader := bufio.NewReader(fp1)
	writer := bufio.NewWriter(fp2)

	buf := make([]byte, BUF_SIZE)

	for {
		n, _ := reader.Read(buf)
		if n == 0 {
			break
		}

		writer.Write(buf[:n])
	}

	fp1.Close()
	fp2.Close()

	elapsed := time.Since(start)
	fmt.Println("The procedure takes ", elapsed)
}
