package main

import "golang.org/x/sync/semaphore"

func main() {
	sw := semaphore.NewWeighted(1)
}
