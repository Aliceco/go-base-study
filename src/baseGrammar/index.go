package main

import "fmt"

//var name string = "An"

//name := "An" // func访问不到

func main() {
	// str :="An" // 定义str并赋值字符串"An"（合适在函数体写法）
	fmt.Println(getMe("An"))
	var name, age = getMe("test")
	fmt.Print(name, age)
}

func getMe(name string) (string, int){
	str := name
	var age int = 12
	return str, age
}