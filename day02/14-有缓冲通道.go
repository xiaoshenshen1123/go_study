package main

import (
	"fmt"
	"time"
)

func main() {
	//有缓冲的管道 numsChan := make(chan int, 10)
	//1-当缓冲写满的时候，写的时候会写阻塞，当被读取后，再恢复写入
	//2-当缓冲区读取完毕，读阻塞
	//3-如果管道没有使用make分配空间,那么管道默认是nil的，读取和写入都会阻塞
	//4-对于一个管道，读和写的次数必须对等
	var names chan string //默认是nil的
	names = make(chan string, 10)
	//names <- "hello"//由于names是nil的，所以写操作会阻塞在这里
	go func() {
		fmt.Println("names:", <-names)
	}()
	names <- "hello"
	time.Sleep(1 * time.Second)

	//读写各50次
	numsChan := make(chan int, 10)
	//写
	go func() {
		for i := 0; i < 50; i++ {
			numsChan <- i
			fmt.Println("写入数据:", i)
		}
	}()
	//读
	func() {
		for i := 0; i < 60; i++ {
			data := <-numsChan
			fmt.Println("读取数据:", data)
		}
	}()
	for {

	}
}
