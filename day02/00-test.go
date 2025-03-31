package main

import "fmt"

func main() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	// 在循环体内通过 value := val 创建新变量 value，
	//每次迭代的 value 都是独立的新变量，地址不同。
	//因此 map 中存储的指针指向正确值。
	for key, val := range slice {
		value := val
		m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k, "===>", *v)
	}
}
