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

	//var me string = "An"
	//var u *string
	//u=&me
	//me = "list"
	//fmt.Println(me, *u)

	var me string = "An"
	fmt.Println(me) // An

	test(&me) // 实参
	//test(me) // 实参

	fmt.Println(me) // abc
}
// 行参
//func test(p string) {
//	p = "abc" // 不会修改me的指针内容（会重新开辟一个指针）
//	fmt.Println(p) // abc
//}

// 行参
func test(p *string) {
	*p = "abc"
}