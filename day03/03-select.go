package main

import (
	"fmt"
	"time"
)

// 当程序有多个channel协同工作，ch1,chan2，某一时刻，ch1和chan2触发了，程序要做响应处理
// 使用select来监听多个通道，当管道被触发时（写入数据、读取数据、关闭管道）
// select语法与switch case很像，但是所有的分支条件都必须是通道io
func main() {
	//var chan1,chan2 chan int
	chan1 := make(chan int)
	chan2 := make(chan int)
	//启动一个go程，负责监听两个channel
	go func() {
		for {
			fmt.Println("监听中.........")
			select {
			case data1 := <-chan1:
				fmt.Println("从chan1读取数据成功.data1:", data1)
			case data2 := <-chan2:
				fmt.Println("从chan2读取数据成功.data2:", data2)
			default:
				fmt.Println("select default 分支called！")
			}
		}
	}()
	//启动go1写chan1
	go func() {
		for i := 0; i < 10; i++ {
			chan1 <- i
			time.Sleep(1 * time.Second / 2)
		}
	}()
	//启动go2写chan2
	go func() {
		for i := 0; i < 10; i++ {
			chan2 <- i
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		fmt.Println("OVER!")
		time.Sleep(5 * time.Second)
	}
}
