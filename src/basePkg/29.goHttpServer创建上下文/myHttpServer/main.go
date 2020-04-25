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

	router.Get("/", func(Context *core.MyContext) {
		Context.WriterString("my string get")
	})

	router.Post("/a", func(Context *core.MyContext) {
		Context.WriterString("my string post")
	})

	http.ListenAndServe(":8099", router)
}
