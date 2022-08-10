package main

import (
	"fmt"
	"sync"
	"time"
)

var msg chan string = make(chan string, 20)

func main() {
	var goroutine_param int = 5 // 协程参数
	var wg sync.WaitGroup
	wg.Add(1) // 主进程需要等待协程结束的数量

	go goroutine_main(goroutine_param, &wg)

	fmt.Printf("goroutine return message: %s\n", <-msg)

	// time.Sleep(10 * time.Second) // 延迟主进程终止时间
	wg.Wait() // 等待协程组结束
	fmt.Println("end of main")
}

func goroutine_main(arg int, wg *sync.WaitGroup) {
	defer wg.Done()

	var cnt int = arg

	for i := 0; i < cnt; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("running goroutine")
	}

	msg <- "Hello, I'm goroutine~"

	fmt.Println("end of goroutine_main")
}
