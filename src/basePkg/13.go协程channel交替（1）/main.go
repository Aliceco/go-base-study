package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func main() {
	users:=strings.Split("shenyi,zhangsan,lisi,wangwu", ",")
	ages:=strings.Split("19,21,25,26", ",")

	//c1:=make(chan bool)
	//c2:=make(chan bool)
	c2,c1:=make(chan bool), make(chan bool)
	ret := make([]string, 0)
	go func() {
		for _, v:= range users {
			<-c1
			ret=append(ret, v)
			c2<-true
		}
		//close(c1)
	}()
	go func() {
		for _, v:= range ages {
			<-c2
			ret=append(ret, v)
			//time.Sleep()
			c1<-true
		}
		//close(c2)
	}()
	c1<-true
	fmt.Println(ret) // [shenyi 19 zhangsan 21 lisi 25 wangwu 26]

	//users:=strings.Split("shenyi,zhangsan,lisi,wangwu", ",")
	//ages:=strings.Split("19,21,25,26", ",")
	//
	//c1:=make(chan string, len(users))
	//c2:=make(chan string, len(ages))
	//
	//go func() {
	//	for _, v:= range users {
	//		c1<-v
	//	}
	//	close(c1)
	//}()
	//go func() {
	//	for _, v:= range ages {
	//		c2<-v
	//	}
	//	close(c2)
	//}()
	//
	//for v:=range c1 {
	//	fmt.Println(v)
	//}
	//for v:=range c2 {
	//	fmt.Println(v)
	//}

	//allList:=make([]string, 0)
	//for i,v:=range users {
	//	allList =append(allList, v)
	//	allList =append(allList, ages[i])
	//}
	//fmt.Println(allList) // [shenyi 19 zhangsan 21 lisi 25 wangwu 26]

	//allList:=make([]string, 0)
	//allList = append(allList, users...)
	//allList = append(allList, ages...)
	//fmt.Println(allList) // [shenyi zhangsan lisi wangwu 19 21 25 26]
}
