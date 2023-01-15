package main

import "fmt"

func main() {
	// 聲明 slice 是一個切片並且初始化, 默認值1,2,3, 長度是3
	//slice1 := []int{1, 2, 3}

	// 宣告 slice 是一個切片, 但是沒有給slice分配空間
	var slice1 []int
	//slice1 = make([]int, 3) //開闢3個空間, 默認值是0

	//宣告 slice 是一個切片,同時給slice分配空間,3個空間,初始化值是0
	//var slice1 []int = make([]int, 3)

	//宣告 slice 是一個切片,同時給slice分配空間,3個空間,初始化值是0, 通過:=推導出slice是一個切片
	//slice1 := make([]int, 3)

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	//判斷一個slice是否為0
	if slice1 == nil {
		fmt.Println("slice 是一個空切片")
	} else {
		fmt.Println("slice 是有空間的")
	}
}
