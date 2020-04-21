package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	//fmt.Println("-----主进程")
	//var wg sync.WaitGroup
	//for i:=1;i<=5;i++{
	//	go func(index int) {
	//		defer func() {
	//			//wg.Add(-1)
	//			wg.Done()
	//		}()
	//		time.Sleep(time.Second*1)
	//		fmt.Println(index, "执行完成")
	//	}(i)
	//	wg.Add(1)
	//}
	//fmt.Println("主进程-----")
	//fmt.Println("任务总数", runtime.NumGoroutine())
	//wg.Wait()
	//fmt.Println("全部执行完成")
	//fmt.Println("-------测试------")
	//-----主进程
	//主进程-----
	//2 执行完成
	//1 执行完成
	//4 执行完成
	//5 执行完成
	//3 执行完成
	//全部执行完成
	//-------测试------


	url:="https://news.cnblogs.com/n/page/%d/"
	var wg sync.WaitGroup

	for i:=1;i<=10;i++ {
		go func(index int) {
			defer func() {
				if err:=recover();err!=nil {
					fmt.Println(err)
				}
				wg.Done()
			}()
			url := fmt.Sprintf(url, index)
			res, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			cnt, _ := ioutil.ReadAll(res.Body)
			ioutil.WriteFile(fmt.Sprintf("./files/%d", index), cnt, 666)
		}(i)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("抓取完毕")

}
