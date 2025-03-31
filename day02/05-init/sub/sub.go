package sub

import "fmt"

// 1-init函数没有参数，没有返回值，原型固定如下
// 2-一个包中包含多个init时，调用顺序时不确定的(同一个包的多个文件都可以有init)
// 3-init函数不允许用户手动调用
// 4-有的时候引用一个包，可能只想使用这个包里面的init函数(mysql的init对驱动进行初始化)
// 但是不想使用这个包里面的其他函数，为了防止编译器报错，可以使用import _"day02/05-init/sub"
// 此时只会调用sub包里的init函数
func init() {
	fmt.Println("this is first init() in package sub ==》 sub.go")
}
func init() {
	fmt.Println("this is second init() in package sub ==》 sub.go")
}

func Sub(a, b int) int {
	// init()不允许显示调用
	return a - b
}
