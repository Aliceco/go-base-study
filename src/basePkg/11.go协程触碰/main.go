package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func sum(max int)  {
	res := 0
	for i:=1; i<=max;i++ {
		res = res+i
	}
	time.Sleep(time.Second*2) // 2秒
	fmt.Println(res)
}
func main() {
	//fmt.Println(((1+100)*100)/2) // 5050
	//sum(100) // 5050
	go sum(100) // 启动线程
	time.Sleep(time.Second*3) // 1秒
}
