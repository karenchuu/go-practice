package main

import (
	"fmt"
)

// select具備多路channel的監控狀態功能
func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			//如果c可寫,則該case就會進來
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		default:
		}
	}

}

func main() {
	c := make(chan int) // 無緩衝的channel
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	//main go
	fibonacii(c, quit)
}
