package main

import "fmt"

func main() {
	names := [7]string{"dock1", "dock2", "dock3", "dock4", "dock5", "dock6", "dock7"}
	// 想基于names创建一个新的数组
	names1 := [3]string{}
	names1[0] = names[0]
	names1[1] = names[1]
	names1[2] = names[2]

	// 切片可以基于一个数组，灵活的创建数组
	names2 := names[0:3] //左闭右开0，1，2要
	fmt.Println(names1, names2)

	names2[1] = "hello"
	fmt.Println("修改names2[1]之后，names2：", names2)
	fmt.Println("修改names2[1]之后，names：", names)
	//1-如果从0元素开始截取，那么冒号左边的数字可以省略
	names3 := names[1:4]
	fmt.Println(names3)
	//2-如果截取到数组的最后一个元素，那么冒号右边的数字可以省略
	names4 := names[5:]
	fmt.Println(names4)
	//3-如果想从左到右全部使用，那么冒号左右两边的数字全省略
	names5 := names[:]
	fmt.Println(names5)
	//4-也可以基于一个字符串进行切片截取，取字符串的子串helloworld
	sub1 := "helloworld"[0:3]
	fmt.Println(sub1) //hel
	//5-可以在创建空切片的时候，明确指定切片的容量，这样可以提供运行效率
	//创建 容量是20 当前长0的string类型切片
	str2 := make([]string, 10, 20) //第三个参数不是必须的，如果没有填写，则默认与长度相同
	fmt.Println("str2len:", len(str2), "str2cap:", cap(str2))
	//6-如果想让切片完全独立于原始的数组，可以使用copy()完成
	namesCopy := make([]string, len(names))
	//names是一个数组，copy函数接收的参数类型是切片，故需要转换为切片
	copy(namesCopy, names[:])
	namesCopy[0] = "hello1"
	fmt.Println("namesCopy", namesCopy, "\nnames:", names)
}
