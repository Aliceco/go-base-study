package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
}
