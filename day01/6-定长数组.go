package main

import "fmt"

func main() {
	// 1.定义一个具有10个数字的数组
	//c语言：int num[10] = {1,2,3,4}
	//go语言: nums := [10]int{1,2,3,4}(常用方式)
	nums := [10]int{1, 2, 3, 4}
	// 2.遍历数组方式1
	for i := 0; i < len(nums); i++ {
		fmt.Printf("i:%d,value:%d\n", i, nums[i])
	}
	// 2.遍历数组方式2 for range
	// key是数组下标 value是数组的值
	for key, value := range nums {
		value += 1
		// value全程只是一个临时变量，不断的被重新赋值，修改它不会修改原始数据
		fmt.Println("i:", key, "value:", value)
	}
	// for range时，如果想要忽略key或者value可以使用_
	//如果两个都忽略那么就不能使用:=，而使用=
	for _, v := range nums {
		fmt.Println("_忽略key，value:", v)
	}
	//不定长数组定义
	//3.使用make创建数组
}
