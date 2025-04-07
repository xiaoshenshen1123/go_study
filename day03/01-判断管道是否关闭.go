package main

import "fmt"

func main() {
	numChan := make(chan int, 10)
	//写
	go func() {
		for i := 0; i < 10; i++ {
			numChan <- i
			fmt.Println("写入数据:", i)
		}
		close(numChan)
	}()
	//读数据
	//for v := range numChan {
	//	fmt.Println("读取数据:", v)
	//}

	for {
		v, ok := <-numChan //ok-idom
		if !ok {
			fmt.Println("管道已经关闭了，准备退出！")
			break
		}
		fmt.Println("v:", v)
	}
	fmt.Println("OVER!")
}
