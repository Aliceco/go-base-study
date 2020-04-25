package main

import (
	"basePkg/myHttpServer/core"
	"net/http"
)
type MyHandler struct {
	
}

func (*MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	//request.URL.Path
	writer.Write([]byte("hello-myServerHttp"))
}
func main() {
	router:=core.DefaultRouter()

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("get"))
	})

	router.Post("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("post"))
	})

	http.ListenAndServe(":8099", router)
}
