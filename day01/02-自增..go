package main

import "fmt"

func main() {
	i := 20
	i++
	//++i也是错误的
	//fmt.Println("i:",i++)// 错误的，i++不允许和其他代码放在一起，必须单独起一行
	fmt.Println("i:", i)
}
