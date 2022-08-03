package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

// 注册子进程函数
func init() {
	Register("childProcess", childProcess)
	if Init() {
		os.Exit(0)
	}
}

// 子进程调用的函数
func childProcess() {
	fmt.Println("Hi, I am a child Process.")
	fmt.Printf("My pid is: %d\n", os.Getpid())
	fmt.Println("End child process.")
}

func main() {
	cmd := Command("childProcess")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil { // 创建子进程
		fmt.Printf("failed to run command: %s", err)
	}

	time.Sleep(30 * time.Second)

	fmt.Println("End parent process.")
}

// 下面是封装的 fork 操作

// -------------------- command_linux --------------------

// Self returns the path to the current process's binary.
// Returns "/proc/self/exe".
func Self() string {
	return "/proc/self/exe"
}

// Command returns *exec.Cmd which has Path as current binary. Also it setting
// SysProcAttr.Pdeathsig to SIGTERM.
// This will use the in-memory version (/proc/self/exe) of the current binary,
// it is thus safe to delete or replace the on-disk binary (os.Args[0]).
func Command(args ...string) *exec.Cmd {
	return &exec.Cmd{
		Path: Self(),
		Args: args,
		SysProcAttr: &syscall.SysProcAttr{
			Pdeathsig: syscall.Signal(0xf),
		},
	}
}

// -------------------- reexec --------------------

var registeredInitializers = make(map[string]func())

// Register adds an initialization func under the specified name
func Register(name string, initializer func()) {
	if _, exists := registeredInitializers[name]; exists {
		panic(fmt.Sprintf("reexec func already registered under name %q", name))
	}

	registeredInitializers[name] = initializer
}

// Init is called as the first part of the exec process and returns true if an
// initialization function was called.
func Init() bool {
	initializer, exists := registeredInitializers[os.Args[0]]
	if exists {
		initializer()

		return true
	}
	return false
}
