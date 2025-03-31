package main

import (
	ADD "day01/12-import/add" // ADD是我们自己重命名的包名
	"day01/12-import/sub"     //sub是文件名也是包名
	"fmt"
)

func main() {
	subResult := sub.Sub(30, 5)  //包名.函数调用
	addResult := ADD.Add(10, 20) //如果一个包里面的函数想对外提供访问权限，函数名必须首字母大写public
	//小写字母开头的函数相当于private，只有相同包名的文件才能使用
	fmt.Println("subResult:", subResult)
	fmt.Println("addResult:", addResult)
}
