package main

import "regexp"

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