package main

import (
	"fmt"
	"net/http"
)
type MyHandler struct {
	
}

func (*MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	
	writer.Write([]byte("hello, myServerHttp"))
}
func main() {
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("hello"))
	//})
	//http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("hello-test"))
	//})
	//err := http.ListenAndServe("127.0.0.1:8099", nil)
	//if err!=nil {
	//	fmt.Println(err)
	//}

	//server:=http.Server{Addr:":8099", Handler:nil}
	//err:=server.ListenAndServe()
	//if err!=nil {
	//	fmt.Println(err)
	//}

	handler:=new(MyHandler)
	server:=http.Server{Addr:":8099", Handler:handler}
	err:=server.ListenAndServe()
	if err!=nil {
		fmt.Println(err)
	}

}
