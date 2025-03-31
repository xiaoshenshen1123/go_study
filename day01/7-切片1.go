package main

import "fmt"

func main() {
	//切片：slice。它的底层也是数组，可以动态改变长度
	//1-定义一个切片，包含多个地名
	names := []string{"duke1", "duke2", "duke3"}
	//2-打印
	for _, value := range names {
		fmt.Println("name:", value)
	}
	fmt.Println("追加元素前的len:", len(names), "追加元素前的容量cap:", cap(names))
	//3-追加数据append之后再赋值给原切片
	names = append(names, "duke4")
	fmt.Println(names)
	//4-对于一个切片，不仅有长度的概念len(),还有容量的概念cap()
	fmt.Println("追加元素后的len:", len(names), "追加元素后的容量cap:", cap(names))
	nums := []int{}
	for i := 0; i < 50; i++ {
		nums = append(nums, i)
		fmt.Println("len:", len(nums), "cap:", cap(nums))
	}
}
