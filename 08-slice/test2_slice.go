package main

import "fmt"

func printArray1(myArray []int) {
	// 引用傳遞, 傳遞的是pointer
	// _ 表示匿名變數 可以不使用
	for _, value := range myArray {
		fmt.Println("value", value)
	}

	myArray[0] = 100
}

func main() {
	myArray := []int{1, 2, 3, 4} // 動態數組, 切片 slice
	fmt.Printf("myArray1 types = %T\n", myArray)
	printArray1(myArray)

	fmt.Println(" ----- ")

	for _, value := range myArray {
		fmt.Println("value", value)
	}
}
