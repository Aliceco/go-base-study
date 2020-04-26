package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, _:=rpc.Dial("tcp", ":8082")
	userName := ""
	err:=client.Call("UserService.GetUserName", 10, &userName)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userName)
}
