package main

import "fmt"

//interface{}是萬用數據類型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	//interface{} 改如何區分 此時引用的底層數據類型到底是什麼

	//給 interface{} 提供 "斷言" 的機制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)
		fmt.Printf("value type is %T]\n", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"Golang"}
	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
}
