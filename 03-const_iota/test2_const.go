package main

import "fmt"

const (
	BEIJING = 10 * iota //iota=0
	SHANGHI             //iota=1
	TAIPEI              //iota=2
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
)

func main() {

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHI = ", SHANGHI)
	fmt.Println("TAIPEI = ", TAIPEI)
}
