package main

import (
	"fmt"
	"time"
)

func main() {
	//-单向读通道
	//var numChanReadOnly <-chan int
	//-单向写通道
	//var numChanWriteOnly chan<- int

	//生产者消费者模型
	//C:数组+锁 thread1：写，thread2：读
	//Go：goroutine+channel
	//1-在主函数中创建一个双向通道
	numChan := make(chan int, 5)
	//2-将numChan传递给producer负责生成
	go producer(numChan) //双向通道可以赋值给同类型的单向通道，单向不能转双向

	//3-将numChan传递给consumer负责消费
	go consumer(numChan)
	time.Sleep(2 * time.Second)
	fmt.Println("OVER!")

}

// producer 生产者 ==>提供一个只写的通道
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
		//data := <-out写通道不允许有读取操作
		fmt.Println("向管道中写入数据:", i)
	}
}

// consumer 消费者 ==>提供一个只读的通道
func consumer(in <-chan int) {
	//in -< 10读通道不允许写数据
	for v := range in {
		fmt.Println("从管道读取数据:", v)
	}
}
