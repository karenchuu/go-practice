package main

import (
	"fmt"
	"time"
)

var ch = make(chan string, 10) // 创建大小为 10 的缓冲 channel

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch <- url // 将 url 发送给 channel
}

func main() {
	for i := 0; i < 3; i++ {
		go download("a.com/" + string(i+'0'))
	}
	for i := 0; i < 3; i++ {
		msg := <-ch // 等待 channel 返回消息
		fmt.Println("finish", msg)
	}
	fmt.Println("Done!")
}
