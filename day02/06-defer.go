package main

import (
	"fmt"
	"os"
)

func main() {
	//1-延迟，关键字，可以用来修饰语句，函数，确保这条语句可以在当前栈退出的时候执行
	//2-一般用于做资源清理工作
	//3-解锁、关闭文件
	//4-在同一个函数中多次调用defer，执行时类似于栈的机制先入后出
	filename := "day02/01-switch.go"
	readFile(filename)
}

func readFile(filename string) {
	//func Open(name string) (*File, error)
	//1-go语言一般会将错误码作为最后一个参数返回
	f1, err := os.Open(filename)
	defer func() {
		fmt.Println("准备关闭文件")
		_ = f1.Close()
	}() // 创建一个匿名函数，同时调用
	//2-err一般nil代表没有错误，执行成功，非nil表示执行失败
	if err != nil {
		fmt.Println("os.Open(\"01-switch.go\")==>打开文件失败:", err)
		return
	}
	defer fmt.Println("000")
	defer fmt.Println("111")
	defer fmt.Println("222")

	//read:func (f *File) Read(b []byte) (n int, err error)
	buf := make([]byte, 1024)
	n, _ := f1.Read(buf)
	fmt.Println("读取文件的实际长度:", n)
	fmt.Println("读取的文件内容:", string(buf))
}
