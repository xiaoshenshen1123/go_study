package main

import "fmt"

func main() {
	// 结构体成员调用时： C语言:ptr->name go语言:ptr.name
	// go语言在使用指针时，会使用内部的垃圾回收机制(gc garbage collector),开发人员不需要手动释放内存
	// c语言不允许返回栈上的指针，go语言可以返回栈上的指针，程序会在编译的时候就确定了变量的分配位置
	// 如果发现有必要的化，就将变量分配到堆上
	name := "lily"
	ptr := &name
	fmt.Println("name:", *ptr)
	fmt.Println("name ptr:", ptr)

	//02-使用new关键字定义一个指针再使用*赋值
	name2Ptr := new(string)
	*name2Ptr = "Duke"
	fmt.Println("name2Ptr:", name2Ptr)
	fmt.Println("name2", *name2Ptr)
	// 可以返回栈上的指针,编译器在编译程序时，会自动判断这段代码，将city变量分配在堆上
	res := testPtr()
	fmt.Println("res city:", *res, "address:", res)
	// 空指针，在c语言:null,在go语言时nil
	//if两端不用加括号
	//if即使有一行代码也要加{}
	if res == nil {
		fmt.Println("res为空， nil")
	} else {
		fmt.Println("res非空， nil")
	}
}

// 定义一个函数，返回string类型的指针
func testPtr() *string {
	city := "深圳"
	ptr := &city
	return ptr

}
