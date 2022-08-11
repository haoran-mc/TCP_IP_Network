package main

import (
	"fmt"
	"sync"

	"github.com/haoran-mc/TCP_IP_Network/ch18/go/semaphore/semaphore"
	// "golang.org/x/sync/semaphore"
)

var sem_one = semaphore.NewSemaphore(1)
var sem_two = semaphore.NewSemaphore(1)
var num int

func main() {
	sem_two.Acquire() // 信号量 sem_two 的初始值设置为 1，为了和 C 代码中保持一致

	var wg sync.WaitGroup
	wg.Add(2)

	go read(&wg)
	go accu(&wg)

	wg.Wait()
}

func read(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Print("Input num: ")

		sem_two.Release() // two--
		fmt.Scanf("%d", &num)
		sem_one.Acquire() // one++
	}
}

func accu(wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0

	for i := 0; i < 5; i++ {
		sem_one.Release() // one--
		sum += num
		sem_two.Acquire() // two++
	}

	fmt.Printf("Result: %d\n", sum)
}
