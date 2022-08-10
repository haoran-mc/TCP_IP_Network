package main

import (
	"fmt"
	"sync"
)

var sum int = 0

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	range1 := []int{1, 5}
	range2 := []int{6, 10}

	go goroutine_summation(range1, &wg)
	go goroutine_summation(range2, &wg)

	wg.Wait()

	fmt.Printf("result: %d\n", sum)
}

func goroutine_summation(area []int, wg *sync.WaitGroup) {
	defer wg.Done()

	start := area[0]
	end := area[1]

	for start <= end {
		sum += start
		start++
	}
}
