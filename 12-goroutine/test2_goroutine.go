package main

import (
	"time"
	"fmt"
)

// 主goroutine
func main() {
	
	/*
	//用go創建乘載一個形參為空，返回值為空的一個函數
	go func() {
		defer fmt.Println("A.defer")
		
		func() {
			defer fmt.Println("B.defer")
			// 退出當前gorutine
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}
	*/

	go func(a int, b, int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		return true
	}(10, 20)

	//死循環
	for {
		time.Sleep(1 * time.Second)
	}
}
