package main

import "fmt"

func printMap(cityMap map[string]string) {
	//cityMap 是一個引用傳遞
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func changeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	//traverse
	printMap(cityMap)

	//刪除
	delete(cityMap, "China")
	//修改
	cityMap["USA"] = "DC"
	changeValue((cityMap))

	fmt.Println(" ---- ")

	//traverse
	printMap(cityMap)
}
