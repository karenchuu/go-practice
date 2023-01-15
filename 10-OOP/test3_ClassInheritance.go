package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

// ====================
type SuperMan struct {
	Human //SuperMan類繼承了Human類的方法

	level int
}

//重定義父類別方法Eat()
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}

func main() {
	h := Human{"Karen", "female"}

	h.Eat()
	h.Walk()

	//定義子類別
	//s := SuperMan{Human{"Li4", "female"}, 88}
	var s SuperMan
	s.name = "li4"
	s.sex = "male"
	s.level = 88

	s.Walk() //父類的方法
	s.Eat()  //子類的方法
	s.Fly()  //子類的方法

	s.Print()
}
