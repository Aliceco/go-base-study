package main

import (
	"fmt"
	"github.com/pquerna/ffjson/ffjson"
)

type NewsModel struct {
	NewsId int
	NewsTitle string
}

func (news NewsModel) ToJSON() string {
	res, err := ffjson.Marshal(news)
	if err != nil {
		return err.Error()
	} else {
		// 返回的[]byte所以转成string
		return string(res)
	}
}
// go get -u github.com/pquerna/ffjson
func main() {
	news := NewsModel{123, "title"}
	//fmt.Println(news.ToJSON())

	var str string = news.ToJSON()
	fmt.Println(str)

	res, _ := ffjson.Marshal(news)
	resDecode := make(map[string]interface{})
	ffjson.Unmarshal(res, &resDecode)
	fmt.Println(resDecode) // map[NewsId:123 NewsTitle:title]
}
