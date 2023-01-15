package main

import (
	"fmt"
)

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

//具體類型
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {
	//b: pair<type:Book, value:book{}地址>
	b := &Book{}

	//r: pair<type: ,value:>
	var r Reader
	//b: pair<type:Book, value:book{}地址>
	r = b

	r.ReadBook()

	var w Writer
	//w: pair<type:Book, value:book{}地址>
	w = r.(Writer) //此處的斷言為什麼會成功? 因為w r具體的type是一致的

	w.WriteBook()
}
