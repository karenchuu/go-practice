package main

import "fmt"

// 本質是一個指針
type AnimalIF interface {
	Sleep()
	GetColor() string //獲取動物的顏色
	GetType() string  //獲取動物的種類
}

// 具體的類
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep() //多態
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())
}

func main() {
	/*
		var animal AnimalIF // 接口的數據類型, 父類指針
		animal = &Cat{"Green"}

		animal.Sleep() // 調用Cat的Sleep(),多態的現象

		animal = &Dog{"Yellow"}
		animal.Sleep() // 調用Dog的Sleep(),多態的現象
	*/

	cat := Cat{"Green"}
	dog := Dog{"Yellow"}

	showAnimal(&cat)
	showAnimal(&dog)
}
