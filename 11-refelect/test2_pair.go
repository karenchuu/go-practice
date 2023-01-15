package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//tty: pair<type:*os.File, value: "/dev/tty"文件描述>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
	}
	//r: pair<type: , value:>
	var r io.Reader
	//r: pair<type:*os.File, value:"/dev/tty">
	r = tty

	//r: pair<type: , value:>
	var w io.Writer
	//w: pair<type:*os.File, value:"/dev/tty">
	w = r.(io.Writer)

	w.Write([]byte("Hello This is A Test!!\n"))
}
