package main

import (
	"fmt"
	"time"
)

func main() {
	var goroutine_param int = 5 // 协程参数

	go goroutine_main(goroutine_param)

	time.Sleep(10 * time.Second) // 延迟主进程终止时间
	fmt.Println("end of main")
}

func goroutine_main(arg int) {
	var cnt int = arg

	for i := 0; i < cnt; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("running goroutine")
	}

	fmt.Println("end of goroutine_main")
}
