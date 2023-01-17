package main

import (
	"fmt"
	"time"
)

func loop(ch chan int) {
	for {
		select {
		case i := <-ch:
			fmt.Println("this  value of unbuffer channel", i)
		}
	}
}

// 缓冲 channel 的阻塞只会发生在 channel 的缓冲使用完的情况下

// fatal error 解决的办法有两个
// 把 channel 开大一点，这是最简单的方法，也是最暴力的
// 把 channel 的信息发送者 ch <- 1 这些代码移动到 go loop(ch) 下面 ，让 channel 实时消费就不会导致阻塞了
func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4 // 这里也会报 fatal error: all goroutines are asleep - deadlock! ，这是因为 channel 的大小为 3 ，而我们要往里面塞 4 个数据，所以就会阻塞住
	go loop(ch)
	time.Sleep(1 * time.Millisecond)
}
