package main

import "fmt"

func main() {

	var a string
	//pair<statictype:string, value:"aceId">
	a = "aceid"

	//pair<type:string, value:"aceId">
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}
