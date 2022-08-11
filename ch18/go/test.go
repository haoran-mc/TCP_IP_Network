package main

import (
	"fmt"
	"time"
)

func test(j int) {
	fmt.Printf("    子子go程%d暂停1s\n", j)
	time.Sleep(time.Second)
	fmt.Printf("    子子go程%d结束\n", j)
}
func main() {
	go func() {
		for j := 0; j < 3; j++ {
			go test(j)
		}
		fmt.Println("  子go程暂停1s")
		time.Sleep(time.Second)
		fmt.Println("  子go程结束")
		// go程及其栈在函数退出时均会销毁
	}()
	println("主程暂停")
	// time.Sleep(time.Second * 5)
	println("主程结束")
}
