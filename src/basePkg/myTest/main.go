package main

import (
	"regexp"
)

func GerRqPath(rq string)string {
	r:=regexp.MustCompile("^GET\\s(.*?)\\sHTTP")
	//r, err:=regexp.Compile(`^GET\\s(.*?)\sHTTP`)
	if r.MatchString(rq) {
		//fmt.Println(r.FindStringSubmatch(rq)) // [GET /abc.php HTTP /abc.php]
		return r.FindStringSubmatch(rq)[1]
	} else {
		return "/"
	}
}

func main() {
	str:=`GET /abc.php HTTP1.1
Host: localhost:8099
Connection: Keep-alive
Upgrade-Insecure-Requests: 1
`
GerRqPath(str)
}
