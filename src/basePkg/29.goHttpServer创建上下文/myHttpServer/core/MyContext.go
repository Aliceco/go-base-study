package core

import "net/http"

type MyContext struct { // 上下文对象
	request *http.Request
	http.ResponseWriter
}

func(this *MyContext) WriterString(str string) {
	this.Write([]byte(str))
}