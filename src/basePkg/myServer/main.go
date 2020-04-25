package main

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8099")
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer lis.Close()
	fmt.Println("创建监听成功，等待客户端连接")

	go func() {
		for {
			fmt.Printf("当前任务数: %d \n", runtime.NumGoroutine())
			time.Sleep(time.Second*2)
		}
	}()

	for {
		client, err := lis.Accept()
		if err!=nil {
			fmt.Println(err)
			return
		}
		go func (c net.Conn) { // 开启协程
			defer c.Close()

			buf:=make([]byte, 4095)
			n, err:= c.Read(buf) // 字节小于或者等于4095直接读出
			if err!=nil {
				if err==io.EOF { // 判断读取完就跳出死循环
					return
				}
				fmt.Println(err)
				return
			}
			if GetRqPath(string(buf[0:n]))=="/delay" {
				time.Sleep(time.Second*5)
			}
			//fmt.Printf(string(buf[0:n]))
			c.Write([]byte(ReadHtml(GetRqPath(string(buf[:n])))))
		}(client)

	}


}
