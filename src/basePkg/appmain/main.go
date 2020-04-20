package main

import (
	"basePkg/servicesb"
	"fmt"
)

func main() {
	//var nservice services.IServer = services.NewService{}
	//fmt.Println(nservice.Get(123)) // 单个新闻内容
	//
	//var uservice services.IServer = services.UserService{}
	//fmt.Println(uservice.Get(123)) // 单个用户信息

	var uservice servicesb.IServer = servicesb.ServiceFactory{}.Create("user")
	fmt.Println(uservice.Get(123)) // 单个用户信息

	var nservice servicesb.IServer = servicesb.ServiceFactory{}.Create("news")
	fmt.Println(nservice.Get(123)) // 单个新闻信息
}
