package main

import "fmt"

func printArray(myArray [4]int) {
	// copy by value
	for index, value := range myArray {
		fmt.Println("index = ", index, "valeu", value)
	}

	myArray[0] = 111
}

func main() {
	var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, "value = ", value)
	}

	//查看數據類型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

	printArray(myArray3)
	fmt.Println(" ----- ")
	for index, value := range myArray3 {
		fmt.Println("index = ", index, "value = ", value)
	}
}
