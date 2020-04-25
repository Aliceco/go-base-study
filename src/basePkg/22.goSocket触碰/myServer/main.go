package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8099")
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer lis.Close()
	fmt.Println("创建监听成功，等待客户端连接")
	client, err := lis.Accept()
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	for {
		// 循环监听怕客户端传过来的数据太多
		buf:=make([]byte, 512)
		n, err:=client.Read(buf) // 字节小于或者等于512直接读出
		if err!=nil {
			if err==io.EOF { // 判断读取完就跳出死循环
				break
			}
			fmt.Println(err)
			return
		}
		fmt.Printf("读取到%d，内容是：%s", n, string(buf))
	}

}
