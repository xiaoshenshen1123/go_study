package main

import "fmt"

// 在c语言里使用typedef int MyInt
type MyInt int

// go语言结构体使用type+struct来处理
type Person struct {
	name   string
	age    int
	gender string
	score  float64
}

func main() {
	var i, j MyInt
	i, j = 1, 2
	fmt.Println("i + j:", i+j)
	// 创建变量并赋值
	lily := Person{
		name:   "lily",
		age:    20,
		gender: "女生",
		score:  80, // 最后一个元素的后面必须加, 如果不加,则必须与}同一行
	}
	fmt.Println("lily:", lily.name, lily.age, lily.gender, lily.score)
	// 结构体没有->操作
	s1 := &lily
	fmt.Println("lily使用指针s1.name打印:", s1.name, s1.age, s1.gender, s1.score)
	fmt.Println("lily使用指针(*s1).name打印:", (*s1).name, (*s1).age, (*s1).gender, (*s1).score)
	// 对结构体赋值时，如果每个字段都赋值了，那么字段名可以省略不写
	// 如果只对部分变量赋值，那么必须明确指定变量名字
	Duke := Person{
		name: "Duke",
		age:  20,
		//gender: "女",
		//score:  80,
	}
	Duke.gender = "男"
	Duke.score = 80
	fmt.Println("Duke:", Duke)
}
