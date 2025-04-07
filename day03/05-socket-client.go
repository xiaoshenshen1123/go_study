package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8848")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	fmt.Println("client与server连接建立成功!")
	sendData := []byte("helloworld")
	for { //保证同一个连接多次向服务器发送数据
		//向服务器发送数据
		cnt, err := conn.Write(sendData)
		if err != nil {
			fmt.Println("conn.Write err:", err)
			return
		}
		fmt.Println("Client===>server cnt:", cnt, "data:", string(sendData))
		//接收服务器返回的数据
		//创建buf用于接收服务器返回的数据
		buf := make([]byte, 1024)
		cnt, err = conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("Client<===Server,cnt:", cnt, "data:", string(buf[:cnt]))
		time.Sleep(1 * time.Second)
	}
	conn.Close()
}
