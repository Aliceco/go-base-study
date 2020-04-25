package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func GetRqPath(rq string)string {
	r:=regexp.MustCompile("^GET\\s(.*?)\\sHTTP")
	//r, err:=regexp.Compile(`^GET\\s(.*?)\sHTTP`)
	if r.MatchString(rq) {
		//fmt.Println(r.FindStringSubmatch(rq)) // [GET /abc.php HTTP /abc.php]
		return r.FindStringSubmatch(rq)[1]
	} else {
		return "/"
	}
}

func Response(body []byte, status HttpStatus) string {
	str:=`HTTP/1.1 %d %s
Server: myServer
Content-type: text/html

%s
`
	ret:=fmt.Sprintf(str, status.Code, status.Message, body)
	return ret
}

func ExisFile(path string) (bool, error){
	_, err := os.Stat("./web"+path)
	if err==nil {
		return true, nil
	} else {
		if os.IsNotExist(err) { // 不存在
			return false, nil
		} else {
			return false, err
		}
	}
}
func ReadHtml(path string) string {
	exist, _:= ExisFile(path)
	if exist {
		file, _ :=ioutil.ReadFile("./web"+path)
		return Response(file, NewHttpStatus(200, "ok"))
	} else {
		return Response([]byte("404"), NewHttpStatus(404, "Not Found"))
	}
}