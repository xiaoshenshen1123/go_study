package main

import (
	"flag"
	"fmt"
	"gin-blog-server/internal/global"
)

func main() {
	//配置文件的路径
	configPath := flag.String("c", "./config.yml", "配置文件的路径")
	flag.Parse()
	//读取配置文件
	conf := global.ReadConfig(*configPath)
	fmt.Println(conf)

}
