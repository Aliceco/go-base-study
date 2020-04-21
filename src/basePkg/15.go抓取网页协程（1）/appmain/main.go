package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
)

func main() {
	
	url:="https://news.cnblogs.com/n/page/%d/"
	c:=make(chan map[int][]byte)

	for i:=1;i<=3;i++ {
		go func(index int) {
			url := fmt.Sprintf(url, index)
			res, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			cnt, _ := ioutil.ReadAll(res.Body)
			c<-map[int][]byte{index:cnt}
			if index==3 {
				close(c)
			}
		}(i)
	}
	for getCnt:=range c {
		for k,v := range getCnt {
			ioutil.WriteFile(fmt.Sprintf("./files/%d", k), v, 666)
		}
	}
}
