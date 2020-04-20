package main

import (
	"basePkg/services"
	"fmt"
)

func main() {
	res := services.GetUser()
	fmt.Println(res)

	//var name string = "An"
	//fmt.Println(name)
	//fmt.Println(&name)
	// &name取地址   *取值

	var name *string // 未初始化指针
	//fmt.Println(name) // <nil>
	name = new(string) // 开辟内存变量并且返回地址
	//fmt.Println(name) // 0xc0000461e0
	*name = "An" // 取值* 赋值*
	fmt.Println(name)
	fmt.Println(*name) // 取值
}

