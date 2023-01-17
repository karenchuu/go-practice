package main

import (
	"fmt"
	"sync"
	"time"
)

/**
Go 语言提供了 sync 和 channel 两种方式支持协程(goroutine)的并发。
例如我们希望并发下载 N 个资源，多个并发协程之间不需要通信，那么就可以使用 sync.WaitGroup，等待所有并发协程执行结束。
*/
var wg sync.WaitGroup

func download(url string) {
	fmt.Println("start to download: ", url)
	time.Sleep(time.Second) // 模拟耗时操作
	wg.Done()
}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)                             // 为 wg 添加一个计数，wg.Done()，减去一个计数
		go download("a.com/" + string(i+'0')) // 启动新的协程并发执行 download 函数
	}
	wg.Wait() // 等待所有的协程执行结束
	fmt.Println("Done!")
}
