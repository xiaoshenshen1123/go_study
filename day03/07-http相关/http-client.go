package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//http包
	client := &http.Client{}
	//func (c *Client) Get(url string) (resp *Response, err error)
	resp, err := client.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("client.Get error:", err)
		return
	}
	//响应包体
	body := resp.Body
	readBodyStr, err := io.ReadAll(body)
	if err != nil {
		fmt.Println("read Body err:", err)
		return
	}
	fmt.Println("body string ", string(readBodyStr))
	//响应头,beego,gin
	ct := resp.Header.Get("Content-Type")
	date := resp.Header.Get("Date")
	server := resp.Header.Get("Server")
	fmt.Println("Content-Type: ", ct)
	fmt.Println("date: ", date)
	fmt.Println("server: ", server)
	//响应行 状态码 url
	url := resp.Request.URL
	code := resp.StatusCode
	status := resp.Status
	fmt.Println("url: ", url)       //http://www.baidu.com
	fmt.Println("code: ", code)     // 200
	fmt.Println("status: ", status) //200 OK
}
