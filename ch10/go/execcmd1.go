package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l", "/var/log")
	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run failed with %s\n", err)
	}
}
