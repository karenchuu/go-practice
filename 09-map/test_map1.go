package main

import "fmt"

func main() {

	// ===> 第一種宣告方式

	var myMap1 map[string]string //一開始宣告的時候為空
	if myMap1 == nil {
		fmt.Println("mayMap1 是一個空map")
	}

	// 在使用map之前, 需要先用make給map分配數據空間
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"

	fmt.Println(myMap1)

	// ===> 第二種宣告方式
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"

	fmt.Println(myMap2)

	// ===> 第三種宣告方式 (一般開發中有初始化就選這個)
	myMap3 := map[string]string{
		"one":   "php",
		"two":   "c++",
		"three": "python",
	}
	fmt.Println(myMap3)
}
