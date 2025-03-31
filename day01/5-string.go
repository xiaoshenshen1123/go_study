package main

import "fmt"

func main() {
	//1.定义
	name := "duke"
	// 需要换行，原生输出字符串时，使用反引号``
	usage := `./a.out <option>
			-h help
			-a xxxx`
	fmt.Println("name:", name)
	fmt.Println("usage:", usage)
	//2.长度，访问
	//c++:name.length
	//GO:自由函数len()
	l1 := len(name)
	fmt.Println("l1:", l1)

	//for循环不需要加()
	for i := 0; i < len(name); i++ {
		fmt.Printf("index:%d,value:%c\n", i, name[i])
	}
	//3.字符串拼接
	i, j := "hello", "world"
	fmt.Println("i+j:", i+j)
}
