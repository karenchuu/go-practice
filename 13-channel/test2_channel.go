package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 3) // 帶有緩衝的channel

	fmt.Println("len(c) = ", len(c), cap(c))

	go func() {
		defer fmt.Println("子go程結束")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("子go程正在運行: ", ",發送的元素=", i, "len(c) = ", len(c), ", cap(c) = ", cap(c))
		}
	}()

	for i := 0; i < 3; i++ {
		num := <-c //從c中接收數據，並賦值給num
		fmt.Println("num = ", num)
	}
	fmt.Println("main結束")
}
