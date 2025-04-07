package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

func main01() {
	//1-使用RPC连接服务器
	//func Dial(network string, address string) (*Client, error)
	//conn, err := rpc.Dial("tcp", "127.0.0.1:8088")
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println("jsonrpc.Dial error:", err)
		return
	}
	defer conn.Close()
	//2-调用远程函数
	//func (client *Client) Call(serviceMethod string, args any, reply any) error
	var reply string //接收函数返回值 ---传出参数
	err = conn.Call("hello.HelloWorld", "李白", &reply)
	if err != nil {
		fmt.Println("Call hello.HelloWorld error:", err)
		return
	}
	fmt.Println(reply)
}

//结合03-design.go测试

func main() {
	Myclient := InitClient("127.0.0.1:8088")
	var reply string
	if err := Myclient.HelloWorld("张三", &reply); err != nil {
		fmt.Println("Call hello.HelloWorld error:", err)
		return
	}
	fmt.Println(reply)
}
