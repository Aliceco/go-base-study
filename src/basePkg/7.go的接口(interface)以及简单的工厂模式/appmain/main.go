package main

import (
	"basePkg/services"
	"fmt"
)

func main() {
	//var nservice services.IServer = services.NewService{}
	//fmt.Println(nservice.Get(123)) // 单个新闻内容
	//
	//var uservice services.IServer = services.UserService{}
	//fmt.Println(uservice.Get(123)) // 单个用户信息

	var uservice services.IServer = services.ServiceFactory{}.Create("user")
	fmt.Println(uservice.Get(123)) // 单个用户信息

	var nservice services.IServer = services.ServiceFactory{}.Create("news")
	fmt.Println(nservice.Get(123)) // 单个新闻信息
}
