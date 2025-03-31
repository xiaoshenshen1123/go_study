package main

import (
	"day02/08-quanxian/src"
	"fmt"
)

func main() {
	s1 := src.Student{
		Hum: src.Human{
			Name:   "张三",
			Age:    20,
			Gender: "男",
		},
		School: "萌宠一中",
	}
	fmt.Println("s1.name:", s1.Hum.Name)
	fmt.Println("s1.school:", s1.School)

	t1 := src.Teacher{}
	t1.Subject = "语文"
	t1.Name = "王老师"
	t1.Gender = "女"
	t1.Age = 35
	fmt.Println("t1:", t1)
	t1.Eat()
	fmt.Println("t1.Human.name", t1.Human.Name)
}
