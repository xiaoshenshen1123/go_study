package main

import "fmt"

// Person类，绑定方法Eat Run Laugh,成员
type Person struct {
	// 成员属性
	name   string
	age    int
	gender string
	score  float64
}

// 在类外面绑定方法
func (this *Person) Eat() {
	//类的方法可以使用自己的成员
	//fmt.Println(this.name + " is eating")
	this.name = "DUKE"
}
func (P Person) Eat2() {
	P.name = "DUKE"
}
func main() {
	lily := Person{
		name:   "Lily",
		age:    20,
		gender: "女",
		score:  100,
	}
	fmt.Println("lily:", lily)
	lily.Eat2()
	fmt.Println("Eat,使用p person不是指针修改姓名", lily)
	lily.Eat()
	fmt.Println("Eat,使用p *person指针修改姓名", lily)
}
