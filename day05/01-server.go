package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

// 定义类对象
type World struct{}

// 绑定类方法
func (this *World) HelloWorld(name string, reply *string) error {
	*reply = name + ",你好!"
	return nil
	//return errors.New("未知的错误!")
}

func main() {
	//1-注册rpc服务，绑定对象方法
	/*func RegisterName(name string, rcvr any) error
	if err := rpc.RegisterName("hello", new(World)); err != nil {
		fmt.Println("注册RPC服务失败！", err)
		return
	}*/
	RegisterService(new(World))
	//2-设置监听
	listener, err := net.Listen("tcp", ":8088")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()
	fmt.Println("开始监听...")
	//3-建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("连接建立成功...")
	//4-绑定服务
	//rpc.ServeConn(conn)
	jsonrpc.ServeConn(conn)
}
