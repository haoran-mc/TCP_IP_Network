package main

import (
	"fmt"
	"sync"
)

const NUM_GOROUTINE int = 100

var num int = 0

func main() {
	var wg sync.WaitGroup
	wg.Add(NUM_GOROUTINE)

	for i := 0; i < NUM_GOROUTINE; i++ {
		if i%2 == 1 {
			go goroutine_inc(&wg)
		} else {
			go goroutine_des(&wg)
		}
	}

	wg.Wait()

	fmt.Printf("result: %d\n", num)
}

func goroutine_inc(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 50000000; i++ {
		num += 1
	}
}

func goroutine_des(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 50000000; i++ {
		num -= 1
	}
}
