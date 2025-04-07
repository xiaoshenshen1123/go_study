package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 要求，服务端在注册rpc服务，绑定对象方法时，能让编译器检测出注册对象是否合法
// 创建接口，在接口中定义方法的原型
type MyInterface interface {
	HelloWorld(name string, reply *string) error
}

// 调用该方法时，需要给i传参，参数i 应该是 实现了HelloWorld方法的类对象  MyInterface
func RegisterService(i MyInterface) {
	rpc.RegisterName("hello", i)
}

//客户端使用
//像调用本地函数一样，调用远程函数

// 定义一个类
type Myclient struct {
	c *rpc.Client
}

// 由于使用了c去调用call，因此需要初始化c
func InitClient(addr string) Myclient {
	conn, _ := jsonrpc.Dial("tcp", addr)
	return Myclient{conn}
}

// 实现函数,原型参照上面的interface实现
func (this *Myclient) HelloWorld(name string, reply *string) error {
	//参数1"hello.HelloWorld"，参照上面的Interface，RegisterName 而来
	return this.c.Call("hello.HelloWorld", name, reply)
}
