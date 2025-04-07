package main

import "fmt"

func main() {
	numsChan2 := make(chan int, 10)
	//写
	go func() {
		for i := 0; i < 50; i++ {
			numsChan2 <- i
			fmt.Println("写入数据:", i)
		}
		fmt.Println("数据全部写完毕，准备关闭管道!")
		close(numsChan2)
	}()
	//遍历管道时只返回一个值，不存在key
	//for range不知道管道是否已经写完了，所以会一直在这里等待
	//解决方法：在写入端，将管道关闭，for range遍历关闭管道时，会退出
	for v := range numsChan2 {
		fmt.Println("读取数据:", v)
	}
	fmt.Println("over!")
}
