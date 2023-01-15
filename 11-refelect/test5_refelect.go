package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", this)
}

func main() {
	user := User{1, "AceId", 18}

	DoFileAndMethod(user)
}

func DoFileAndMethod(input interface{}) {
	//獲取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is:", inputType.Name())

	//獲取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is:", inputValue)

	//通過type 獲取裡面的欄位
	//1. 獲取interface的reflect.Type,通過Type獲得NumField,進行便利
	//2. 得到每個field, 數據類型
	//3. 通過field有一個Interface()方法得到對應的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	//通過type 獲取裡面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
