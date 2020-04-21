package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

func main() {
	//list:=make([]int, 0)// 空切片
	//for i:=1; i<=10;i++ {
	//	list=append(list, 1)
	//}
	//fmt.Println(len(list))

	list:=make([]int, 0)// 空切片
	var wg sync.WaitGroup
	var mutex sync.Mutex

	go func() {
		defer wg.Done()
		defer mutex.Unlock() // 解锁
		mutex.Lock() //上锁
		for i:=1; i<=1000000;i++ { // 执行到一定时间-如果还没有执行完成-协程调度-会自动把cpu让出
			list=append(list, 1)
		}
		fmt.Println(len(list))

	}()

	go func() {
		defer wg.Done()
		defer mutex.Unlock() // 解锁
		mutex.Lock() //上锁
		for i:=1; i<=1000000;i++ {
			list=append(list, 1)
		}
		fmt.Println(len(list))
	}()
	wg.Add(2)
	wg.Wait()
	fmt.Println(len(list))
}
