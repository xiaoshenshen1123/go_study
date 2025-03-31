package main

import "fmt"

// 1-函数返回值在参数列表之后
// 2-如果有多个返回值，需要使用圆括号包裹，多个参数之间使用,分割
func test1(a int, b int, c string) (int, string, bool) {
	return a + b, c, true
}

// 返回值可以有名字，可以直接简写return
func test2(a, b int, c string) (res int, str string, b1 bool) {
	res = a + b
	str = c
	b1 = true
	return
}

// 返回值只有一个参数，并且没有名字，那么不需要加()
func test3(a, b int) int {
	return a + b
}
func main() {
	v1, s1, _ := test1(10, 20, "hello")
	fmt.Println(v1, s1)
	v2, s2, q2 := test2(10, 20, "hello")
	fmt.Println(v2, s2, q2)
	v3 := test3(20, 20)
	fmt.Println(v3)
}
