package main

import (
	"fmt"
	"time"
)

func loop(ch chan int) {
	// 这里会报错 fatal error: all goroutines are asleep - deadlock! 就是因为 ch<-1 发送了，但是同时没有接收者，所以就发生了阻塞
	for {
		select {
		case i := <-ch:
			fmt.Println("this  value of unbuffer channel", i)
		}
	}
}

func main() {
	ch := make(chan int)
	ch <- 1 // 但如果我们把 ch <- 1 放到 go loop(ch) 下面，程序就会正常运行
	go loop(ch)
	time.Sleep(1 * time.Millisecond)
}
