package main

import (
	"net/http"
	"time"
)
type MyHandler struct {
	
}

func (*MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	//request.URL.Path
	writer.Write([]byte("hello-myServerHttp"))
}
func main() {
	//myMux:=http.NewServeMux()
	//myMux.Handler("/", &MyHandler{})
	//http.ListenAndServe(":8099", myMux)

	myMux:=http.NewServeMux()
	myMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("index"))
	})
	myMux.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		getUserName:=request.URL.Query().Get("userName")
		if getUserName!="" {
			c:=&http.Cookie{Name:"userName", Value:getUserName, Path:"/"}
			http.SetCookie(writer, c)
		}
		writer.Write([]byte("登录页面"))
	})
	myMux.HandleFunc("/logout", func(writer http.ResponseWriter, request *http.Request) {
		c:=&http.Cookie{Name:"userName", Value:"", Path:"/", Expires:time.Now().AddDate(-1, 0, 0)}
		http.SetCookie(writer, c)
		writer.Header().Set("Content-type", "text/html")
		writer.Write([]byte("logout del cookie...."))
		script:=`<script>
			setTimeout(()=> {
			 self.location='/'
			}, 2000)
</script>`
		writer.Write([]byte(script))
	})
	http.ListenAndServe(":8099", myMux)
}
