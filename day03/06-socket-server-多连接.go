package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	ip := "127.0.0.1"
	port := 8848
	address := fmt.Sprintf("%s:%d", ip, port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("net.listen err:", err)
	}

	//需求:server可以接收多个连接==>主go程负责监听,子go程负责数据处理
	//每个连接可以接收处理多轮数据请求
	for {
		fmt.Println("监听中...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.accept err:", err)
			return
		}
		fmt.Println("连接建立成功！")
		go handleFunc(conn)
	}
}

// 处理具体业务的逻辑，需要将conn传递进来，每一个新连接，conn是彼此独立的
func handleFunc(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		fmt.Println("准备读取客户端发送的数据")
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("Client ===> Server,长度:", cnt, "数据内容:", string(buf))
		upperData := strings.ToUpper(string(buf[:cnt]))
		cnt, err = conn.Write([]byte(upperData))
		fmt.Println("Client<=====Server,长度:", cnt, "数据内容:", upperData)
	}
	//_ = conn.Close()
}
