package main

import "fmt"

// 如果類名首字母大小, 表示其他包也能夠訪問
type Hero struct {
	//如果類屬性首字母大小, 表示該屬性是對外訪問,否則只能夠類的內部訪問
	Name  string
	Ad    int
	Level int
}

/* 值傳遞
func (this Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

// this就綁定到當前的sturct
func (this Hero) GetName() string {
	fmt.Println("Name = ", this.Name)
	return this.Name
}

func (this Hero) SetName(newName string) {
	//this是調用該方法的對象的一個副本拷貝
	this.Name = newName
}
*/

/**
通常定義對象用引用傳遞 需要加上*
*/
func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

// this就綁定到當前的sturct
func (this *Hero) GetName() string {
	fmt.Println("Name = ", this.Name)
	return this.Name
}

func (this *Hero) SetName(newName string) {
	//this是調用該方法的對象的一個副本拷貝
	this.Name = newName
}

func main() {
	hero := Hero{Name: "Karen", Ad: 100, Level: 1}
	hero.Show()

	hero.SetName("li4")

	hero.Show()
}
