package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
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

	url:="https://news.cnblogs.com/n/page/%d/"
	c:=make(chan map[int][]byte)

	for i:=1;i<=3;i++ {
		go func(index int) {
			defer func() {
				if err:=recover();err!=nil {
					fmt.Println(err)
					if index==3 {
						close(c)
					}
				}
			}()
			url := fmt.Sprintf(url, index)
			res, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			cnt, _ := ioutil.ReadAll(res.Body)
			if index==3 {
				test()
			}
			c<-map[int][]byte{index:cnt}



		}(i)
	}
	for getCnt:=range c {
		for k,v := range getCnt {
			ioutil.WriteFile(fmt.Sprintf("./files/%d", k), v, 666)
		}
	}
}
