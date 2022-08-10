package main

import (
	"fmt"
	"sync"
)

const NUM_THREAD int = 100

var num int = 0
var mutex sync.Mutex //

func main() {
	var wg sync.WaitGroup
	wg.Add(NUM_THREAD)

	for i := 0; i < NUM_THREAD; i++ {
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

	mutex.Lock() // 上锁
	for i := 0; i < 50000000; i++ {
		num += 1
	}
	mutex.Unlock() // 解锁
}

func goroutine_des(wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	for i := 0; i < 50000000; i++ {
		num -= 1
	}
	mutex.Unlock()
}
