package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type UserService struct {
}
func (this *UserService) GetUserName(userId int, userName *string) error{
	if userId==101{
		*userName="An"
	} else {
		*userName="guest"
	}
	return nil
}


func main() {
	list, _:=net.Listen("tcp", ":8082") // 注册监听
	rpc.Register(new(UserService)) // 注册对外暴露的方法
	for {
		client,_:=list.Accept() // 获取连接
		//rpc.ServeConn(client) // 处理客户端连接
		jsonrpc.ServeConn(client) // 处理客户端连接
	}
}
