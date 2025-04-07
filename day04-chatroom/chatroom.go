package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type User struct {
	Name string
	Id   string
	msg  chan string
}

// 创建一个全局的map结构，用于保存所有用户
var allUsers = make(map[string]User)

// 定义一个公共的message通道，用于接收任何人发送的消息
var messages = make(chan string, 10)

func main() {
	//创建服务器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.listen err:", err)
		return
	}
	//启动全局唯一的go程负责监听message通道，写给所有的用户msg通道
	go broadcast()
	fmt.Println("服务器启动成功")
	//想要能够建立多个连接，在这里加for
	for {
		fmt.Println("11111主go程监听中....")
		//监听
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.accept err:", err)
			return
		}
		//建立连接
		fmt.Println("建立连接成功！")
		//启动处理业务的go程
		go handler(conn)
	}
}

// 处理具体业务
func handler(conn net.Conn) {
	fmt.Println("启动业务....")
	//客户端与服务器建立连接时,会有ip和port==》当成user的id
	clientAddr := conn.RemoteAddr().String()
	fmt.Println("clientAddr:", clientAddr)
	//创建user
	newUser := User{
		Name: clientAddr,            //可以修改，提供rename命令修改，建立连接时默认初始值与id相同
		Id:   clientAddr,            //id不会修改，这个作为在map中的key
		msg:  make(chan string, 10), //注意需要make空间，否则无法写入数据
	}
	//添加user到map结构
	allUsers[newUser.Id] = newUser
	// 定义一个退出信号，用于监听client退出
	var isQuit = make(chan bool, 1)
	//创建一个用于重置计数器的管道，用于告知watch函数，当前用户正在输入
	var resetTimer = make(chan bool, 1)
	//启动go程，负责监听退出信号
	go watch(&newUser, conn, isQuit, resetTimer)
	//启动go程，负责将msg信息返回客户端
	go writeBackToClient(&newUser, conn)
	//向message写入数据，当前用户上线的信息，用于通知所有人(广播)
	loginInfo := fmt.Sprintf("[%s]:[%s] ===>上线了,login!\n", newUser.Id, newUser.Name)
	messages <- loginInfo
	for {
		//具体业务逻辑
		buf := make([]byte, 1024)
		//读取客户端发送来的请求数据
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("客户端主动关闭ctrl+c，准备退出！")
			//map删除用户，conn close掉
			//服务器还可以主动退出
			//在这里不进行真正的退出动作，而是发送一个退出信号，统一做退出处理，可以使用新的管道来做信号
			isQuit <- true
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("服务器读取到来自client的数据是:", string(buf[:n-1]), "长度为:", n)
		//--------业务逻辑处理 开始-------
		//1-查询当前所有用户 who
		// a.先判断接收的数据是不是who==》长度&&字符串
		userInput := string(buf[:n-1]) //这是用户输入的数据，最后一个是回车，去掉这个长度
		if len(userInput) == 4 && userInput == "\\who" {
			// b.遍历allUser这个map==>key:=userid value:user本身,把id和name拼接成一个字符串返回给客户端
			fmt.Println("用户即将查询所有用户信息!")
			var userInfos []string //这个切片包含所有的用户信息
			//[]string{"id:1,name:小明","id:1,name:小明","id:1,name:小明"}
			for _, user := range allUsers {
				userInfo := fmt.Sprintf("userid:%s,username:%s", user.Id, user.Name)
				userInfos = append(userInfos, userInfo)
			}
			//最终写到管道中，一定是一个字符串，而不是切片数组
			r := strings.Join(userInfos, "\n")
			//将数据返回给查询的客户端
			newUser.msg <- r
		} else if len(userInput) > 9 && userInput[:7] == "\\rename" {
			//1.读取数据判断长度7,判断字符是rename
			//2.使用|进行分割，获取|后的部分作为名字
			strs := strings.Split(userInput, "|")
			//3.更新用户名字newUser.name = Duke
			newUser.Name = strs[1]
			allUsers[newUser.Id] = newUser //更新map中的user
			//4.通知客户端，更新成功
			newUser.msg <- "rename更新成功!"
		} else {
			//如果用户输入的不是命令，只是普通的聊天信息，那么只需要写到广播通道中即可，由其他的go程进行常规转发
			messages <- string(userInput)
		}
		resetTimer <- true
		//--------业务逻辑处理 结束-------

	}
}

// 向所有用户广播消息，启动全局唯一的go程
func broadcast() {
	fmt.Println("广播go程启动成功...")
	defer fmt.Println("broadcast 程序退出!")
	for {
		//1-从message管道中读取数据
		fmt.Println("broadcast监听message管道中...")
		info := <-messages
		fmt.Println("message接收到消息:", info, "将要传递给user.msg管道中")
		//2-将数据写到每一个用户的msg管道中
		for _, user := range allUsers {
			//在这里阻塞了，msg是非缓冲的,需要改成缓冲的
			user.msg <- info
		}
	}
}

// 每个用户应该还有一个用来监听自己msg管道的go程,负责将数据返回客户端
func writeBackToClient(user *User, conn net.Conn) {
	//不断读取自己的msg管道
	fmt.Printf("22222 user:%s的go程正在监听自己的msg管道\n", user.Name)
	for data := range user.msg {
		fmt.Printf("user:%s 写回客户端的数据为:%s \n", user.Name, data)
		//写回客户端
		_, _ = conn.Write([]byte(data))
	}
}

// 启动一个go程用于监听退出信号，触发后进行清理工作:delete map,close conn
func watch(user *User, conn net.Conn, isQuit <-chan bool, resetTimer <-chan bool) {
	fmt.Println("33333 启动监听退出信号的go程....")
	defer fmt.Println("watch go程退出！")
	for {
		select {
		case <-isQuit:
			logoutInfo := fmt.Sprintf("用户:%s 主动退出了！\n", user.Name)
			fmt.Println("删除当前用户:", user.Name)
			delete(allUsers, user.Id)
			messages <- logoutInfo
			conn.Close()
			return
		case <-time.After(10 * time.Second):
			logoutInfo := fmt.Sprintf("用户:%s 超时退出了！\n", user.Name)
			fmt.Println("删除当前用户:", user.Name)
			delete(allUsers, user.Id)
			messages <- logoutInfo
			conn.Close()
			return
		case <-resetTimer:
			fmt.Printf("连接%s 重置计数器！\n", user.Name)
		}
	}
}
