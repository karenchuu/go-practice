package main

import "fmt"

/*
func swap(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}
*/

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println("a = ", a, "b = ", b)

	var p *int
	p = &a // p = a的地址

	fmt.Println(&a)
	fmt.Println(p)

	var pp **int // 二級指針: 一級指針的地址
	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)
}
