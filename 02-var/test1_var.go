package main

import (
	"fmt"
)

var gA int = 100
var gB = 200

//gc := 200

func main() {
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Println("b = ", b)
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Println("b = ", b)
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	// 不支持全局變數
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := "100"
	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	fmt.Println("gA = ", gA, "gB = ", gB)
	fmt.Println("gA = ", gA, "gB = ", gB)

	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, "yy = ", yy)

	var kk, ll = 100, "ACE"
	fmt.Println("kk = ", kk, "ll = ", ll)

	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, "jj = ", jj)
}
