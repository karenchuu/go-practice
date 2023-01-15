package main

import "fmt"

func main() {

	// defer是壓棧的方式, 所以end2會先執行
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main: hello go 1")
	fmt.Println("main: hello go 2")
}
