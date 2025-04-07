package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	Name    string `json:"-"` //==>在使用json编码时，这个编码不参与
	Subject string `json:"subject"`
	Age     int    `json:"age,string"`        //在json编码时，将age转成string类型，一定要是两个字段:名字,类型，中间不能加空格
	Address string `json:"address,omitempty"` //在json编码时，如果这个字段是空的，那么不参与编码
}

func main() {
	t1 := Teacher{
		Name:    "duke",
		Subject: "语文",
		Age:     28,
	}
	fmt.Println("t1:", t1)
	encodeInfo, err := json.Marshal(t1)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	fmt.Println("编码后的t1:", string(encodeInfo))
}
