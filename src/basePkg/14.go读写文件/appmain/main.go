package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

func main() {
	// 写入文件
	//content:="hello, An"
	//ioutil.WriteFile("./files/1.txt", []byte(content), 666)
	// 读取文件
	getFile,_ := ioutil.ReadFile("./files/1.txt")
	fmt.Println(string(getFile))

	//url:="https://news.cnblogs.com/n/page/1/"
	//res,err:=http.Get(url)
	//if err!=nil {
	//	fmt.Println(err)
	//}
}
