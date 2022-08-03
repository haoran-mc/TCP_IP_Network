package main

import (
	"fmt"
	"os"

	"github.com/haoran-mc/TCP_IP_Network/ch10/go/fork/reexec"
)

func init() {
	fmt.Printf("init start, os.Args = %+v\n", os.Args)
	reexec.Register("childProcess", childProcess)
	if reexec.Init() {
		os.Exit(0)
	}
}

func childProcess() {
	fmt.Println("childProcess")
}

func main() {
	fmt.Printf("main start, os.Args = %+v\n", os.Args)
	cmd := reexec.Command("childProcess")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Printf("failed to run command: %s", err)
	}
	if err := cmd.Wait(); err != nil {
		fmt.Printf("failed to wait command: %s", err)
	}

	fmt.Println("main exit")
}
