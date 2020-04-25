package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net"
)

func main() {
	conn, err:=net.Dial("tcp", "127.0.0.1:8099")
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	conn.Write([]byte("i am an"))
	buf:=make([]byte, 4095)
	n, err:=conn.Read(buf)
	if err!=nil {
		if err==io.EOF { // 判断读取完就跳出死循环
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf(string(buf[0:n]))
}
