package main

import "fmt"

func main() {
	// 变量定义:var
	// 常量定义:const
	//01-先定义变量再赋值 var 变量名 数据类型
	var name string
	name = "duke" //ctrl + alt + l可以快速格式化代码
	fmt.Println("name:", name)
	var age int
	age = 20
	fmt.Printf("name:%s,age:%d", name, age)
	//02-定义变量时直接赋值 var 变量名 数据类型 = 值
	var gender string = "男"
	fmt.Println("gender:", gender)
	//03-定义直接赋值，使用自动推导(最常用)
	address := "北京"
	fmt.Println("address:", address)
	//灰色部分表示形参
	test(10, "str")
	//04-平行赋值
	i, j := 10, 20 // 同时定义两个变量
	fmt.Println("变换前i:", i, "变换前j:", j)
	// 互换
	i, j = j, i
	fmt.Println("变换后i:", i, "变换后j:", j)
}
func test(a int, b string) {
	fmt.Println(a)
	fmt.Println(b)
}
