package main

import "fmt"

func main() {
	// 标签 LABEL1
	// continue LABEL1==》会跳到指定的位置，但是也会记录之前的状态，i变成1
	// goto LABEL1==》下次进入循环时，i不会保存之前的状态，而是重新从0开始计算
	// break LABEL1==》直接跳出指定位置的循环
LABEL1:
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j == 3 {
				goto LABEL1
				//continue LABEL1
				//break LABEL1
			}
			fmt.Println("i:", i, "j:", j)
		}
	}
	fmt.Println("over!")
}
