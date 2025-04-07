package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//func ListenAndServe(addr string, handler Handler) error
	//常规写法
	/*err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("http server err:", err)
		return
	}*/

	//注册路由 router
	//xxx/user ===> func1
	//xxx/name ===> func2
	//xxx/id ==> func3
	//https://127.0.0.1:8080/user,func是回调函数，用于路由的响应，函数原型是固定的
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		//request==》包含客户端发来的数据
		fmt.Println("用户请求详情:")
		fmt.Println("request:", request)
		//writer==》通过writer将数据返回给客户端
		_, _ = io.WriteString(writer, "这是/user请求返回的数据")
	})
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是/name请求返回的数据")
	})
	http.HandleFunc("/id", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是/id请求返回的数据")
	})
	fmt.Println("httpServe start ...")
	//精简写法
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println("http server err:", err)
		return
	}

}
