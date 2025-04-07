package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id     int
	Name   string
	Age    int
	Gender string //如果是小写gender那么序列化编码不成功
}

func main() {
	lily := Student{
		Id:     1,
		Name:   "lily",
		Age:    20,
		Gender: "male",
	}
	//编码(序列化)==》把结构变成字符串
	//func Marshal(v any) ([]byte, error)
	encodeInfo, err := json.Marshal(lily)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
	}
	fmt.Println(string(encodeInfo))

	//对端接收到数据
	//反序列化
	var lily2 Student
	if err := json.Unmarshal(encodeInfo, &lily2); err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	fmt.Println("id:", lily2.Id)
	fmt.Println("name:", lily2.Name)
	fmt.Println("age:", lily2.Age)
	fmt.Println("gender:", lily2.Gender)
}
