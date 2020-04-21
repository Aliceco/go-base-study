package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func sum(min, max int, c chan int) {
	res := 0
	for i:=min; i<=max;i++ {
		res = res+i
	}
	c <- res // 发送数据到c
}
func main() {
	// 同步执行 1.0999511s 500000000500000000
	// 单个协程 1.2963141s 500000000500000000
	// 拆成两个协程 420.6758ms 125000000250000000 375000000250000000 500000000500000000
	num:=1000000000
	start:=time.Now()

	c:=make(chan int, 2) // 2个容量
	go sum(1,num/2, c) // 启动协程
	go sum(num/2+1, num, c) // 启动协程

	ret1, ret2 := <- c, <- c // 读取数据，接收者
	//ret1:= <- c // 读取数据，接收者
	//ret2 := <- c // 读取数据，接收者

	end:=time.Now()

	fmt.Println(end.Sub(start), ret1, ret2, ret1+ret2)
}

//func sum(max int, c chan int) {
//	res := 0
//	for i:=1; i<=max;i++ {
//		res = res+i
//	}
//	c <- res // 发送数据到c
//}
//func main() {
//	c:=make(chan int, 2) // 2个容量
//	go sum(1000000, c) // 启动线程
//	c <- 123 // 发送数据到C
//	ret := <- c // 读取数据，接收者
//	fmt.Println(ret)
//}

//func main() {
//	c:=make(chan int)
//	go sum(100, c) // 启动线程
//	ret := <- c // 读取数据，接收者
//	fmt.Println(ret)
//}
//
//func sum(max int, c chan int) {
//	res := 0
//	for i:=1; i<=max;i++ {
//		res = res+i
//	}
//	c <- res // 发送数据到c
//}
