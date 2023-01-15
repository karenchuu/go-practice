package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	//[0, 2]
	s1 := s[0:2] //[1, 2]

	fmt.Println(s1)

	s1[0] = 100

	fmt.Println(s)
	fmt.Println(s1)

	//copy 可以將底層的slice一起進行拷貝
	s2 := make([]int, 3) // s2 = [0,0,0]

	//將s中的值一次拷貝到s2中
	copy(s2, s)
	fmt.Println(s2)
}
