package main

import (
	"fmt"
)

func main() {
	c := make(chan int) // 無緩衝的channel

	go func() {
		for i := 0; i < 3; i++ {
			c <- i
		}

		//可以關閉一個channel
		close(c)
	}()

	for i := 0; i < 3; i++ {
		//ok如為true表示channel沒關閉, 如果為false表示channel已經關閉
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("Main Finished..")
}
