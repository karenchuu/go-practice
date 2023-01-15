package main

import (
	"fmt"
)

func main() {
	//定義一個channel
	c := make(chan int)

	go func() {
		defer fmt.Println("子goroutine end")

		fmt.Println("子goroutine running...")

		c <- 666 //將666發送給c
	}()

	num := <-c // 從c中接收數據，並賦值給num

	fmt.Println("num = ", num)
	fmt.Println("main goroutine end...")
}
