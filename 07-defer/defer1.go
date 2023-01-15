package main

import "fmt"

func deferA() {
	fmt.Println("A")
}

func deferB() {
	fmt.Println("B")
}

func deferC() {
	fmt.Println("C")
}

func main() {

	deferA()
	deferB()
	deferC()
}
