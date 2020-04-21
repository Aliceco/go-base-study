package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"time"
)

//func test()  {
//	a:=1
//	defer func() {
//		a++
//		fmt.Println(a)
//	}()
//	panic("my exp")
//	a++
//	fmt.Println(a)
//}
//func test()  {
//	a:=1
//	// 后进先出
//	defer func() {
//		a++
//		fmt.Println("----a---")
//	}()
//	defer func() {
//		fmt.Println(a)
//		fmt.Println("----b---")
//	}()
//	panic("my exp")
//	a++
//	fmt.Println(a)
//}

// 返回1
//func test() int {
//	a:=1
//	defer func() {
//		a++
//	}()
//	return a
//}

// 返回2
/**
return a会先把a赋值给ret    ret就是a   a++
*/
//func test() (ret int) {
//	a:=1
//	defer func() {
//		a++
//	}()
//	return a
//}
// 返回10
//func test() (ret int) {
//	defer func() {
//		ret=10
//	}()
//	return
//}


//func test()(ret int) {
//	defer func() {
//		ret = -1
//	}()
//	panic("exp")
//}
//func main()  {
//	defer func() {
//		err:=recover()
//		if err!=nil {
//			fmt.Println(err)
//		}
//	}()
//	fmt.Println("----a----")
//	fmt.Println(test())
//	fmt.Println("----b----")
//}


func test() {
	panic("exp")
}
func main() {
	//c:make(chan int) // 读写
	//c:make(<-chan int) // 只读
	//c:make(chan<-int) // 只写

	//c:time.After(time.Second*3)
	//fmt.Println("不阻塞")
	//fmt.Println(<-c) // 3秒后出现结果

	// label(标签)语法
	// 死循环
	//i:=0
	//abc:fmt.Println(i) // 1 , 2, 3, 4 ,5 .......
	//time.Sleep(time.Second*2)
	//i++
	//goto abc

	url:="https://news.cnblogs.com/n/page/%d/"
	c:=make(chan map[int][]byte)

	for i:=1;i<=3;i++ {
		go func(index int) {
			defer func() {
				if err:=recover();err!=nil {
					fmt.Println(err)
					close(c)
				}
			}()
			url := fmt.Sprintf(url, index)
			res, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			cnt, _ := ioutil.ReadAll(res.Body)

			c<-map[int][]byte{index:cnt}
		}(i)
	}
	//i:=0
	//for getCnt:=range c { // 主线程干的事情，不断的range channel，直到channel被关闭
	//	i++
	//	for k,v := range getCnt {
	//		ioutil.WriteFile(fmt.Sprintf("./files/%d", k), v, 666)
	//	}
	//	if (i==3) {
	//		close(c)
	//	}
	//}
	res := map[int][]byte{}
	myloop:for { // 死循环
		select {
			case res=<-c:
				for k,v := range res {
					ioutil.WriteFile(fmt.Sprintf("./files/%d", k), v, 666)
				}
			case <-time.After(time.Second*3):
				break myloop
		}
	}

}
