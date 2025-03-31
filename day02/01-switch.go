package main

import (
	"fmt"
	"os"
)

// 从命令行中输入参数，在switch中进行处理
func main() {
	// GO：os.Args ==> 直接可以获取命令输入，是一个字符串切片
	cmds := os.Args
	//os.Args[0] ==>是程序名称
	//os.Args[1] ==>是命令行输入的第一个参数
	for key, cmd := range cmds {
		fmt.Println("key:", key, "cmd:", cmd, "cmdlen:", len(cmd))
	}
	if len(os.Args) < 2 {
		fmt.Println("请输入正确的参数")
	}
	switch cmds[1] {
	case "hello":
		fmt.Println("hello")
		// go的switch默认加上了break，不需要手动处理
		// 如果想向下穿透需要加关键字fallthrough
		fallthrough
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("default")
	}
}
