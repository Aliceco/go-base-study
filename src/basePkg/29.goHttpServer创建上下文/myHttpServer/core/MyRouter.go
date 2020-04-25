package core

import "net/http"

type MyHandlerFunc func(Context *MyContext)

type MyRouter struct {
	Mapping map[string]map[string]MyHandlerFunc
	Context *MyContext
}

func DefaultRouter() *MyRouter{
	return &MyRouter{make(map[string]map[string]MyHandlerFunc), &MyContext{}}
}

func (this MyRouter) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	this.Context.request = request
	this.Context.ResponseWriter = writer

	// chrome 会默认请求图标地址
	if f, OK:=this.Mapping[request.Method][request.URL.Path]; OK{
		f(this.Context)
	}
}

func (this MyRouter) Get(path string, f MyHandlerFunc)  {
	if this.Mapping["GET"]==nil {
		this.Mapping["GET"]=make(map[string]MyHandlerFunc)
	}
	this.Mapping["GET"][path] = f
}

func (this MyRouter) Post(path string, f MyHandlerFunc)  {
	if this.Mapping["POST"]==nil {
		this.Mapping["POST"]=make(map[string]MyHandlerFunc)
	}
	this.Mapping["POST"][path] = f
}