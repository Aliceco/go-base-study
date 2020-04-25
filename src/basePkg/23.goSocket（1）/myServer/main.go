package main

import (
	"fmt"
	"io"
	"net"
)
func res() string {
	str:=`HTTP/1.1 200 OK
Server: myServer
Content-type: text/html

this is body
`
	return str
}
func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8099")
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer lis.Close()
	fmt.Println("创建监听成功，等待客户端连接")

	for {
		client, err := lis.Accept()
		if err!=nil {
			fmt.Println(err)
			return
		}
		func (c net.Conn) {
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
			fmt.Printf(string(buf[0:n]))
			c.Write([]byte("测试abc"))
		}(client)

	}


}
