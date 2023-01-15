package main

import "fmt"

//聲明一種新的數據類型, 是int的一個別名
type myint int

//定義一個結構體, 把多種數據類型或屬性結合在一起 跟Class一樣
type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	//傳遞一個book副本
	book.auth = "666"
}

func changeBook2(book *Book) {
	//傳遞一個book副本
	book.auth = "777"
}

func main() {
	/* 不是重點
	var a myint = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T", a)
	*/

	var book1 Book
	book1.title = "Golang"
	book1.auth = "Karen"

	fmt.Printf("%v\n", book1)

	changeBook(book1)

	fmt.Printf("%v\n", book1)

	changeBook2(&book1)

	fmt.Printf("%v\n", book1)
}
