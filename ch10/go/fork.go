package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func init() {
	Register("childProcess", childProcess)
	if Init() {
		os.Exit(0)
	}
}

func childProcess() {
	fmt.Println("child process")
	fmt.Println(os.Getpid())
}

var gval int = 10

func main() {
	var lval int = 20
	gval++
	lval += 5

	cmd := Command("childProcess")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Printf("failed to run command: %s", err)
	}
	fmt.Println(os.Getpid())
	if err := cmd.Wait(); err != nil {
		fmt.Printf("failed to wait command: %s", err)
	}
}

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
