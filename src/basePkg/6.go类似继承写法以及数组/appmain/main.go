package main

import (
	"basePkg/submodels"
	"fmt"
)

func main() {
	var sn submodels.SportsNew
	//sn.News.Uid
	//sn.Uid
	sn.Uid = 133
	sn.Uname = "An"
	fmt.Println(sn.ToString()) // 用户名是An

	sn.Tags=[]string{"足球", "篮球", "排球"}
	fmt.Println(sn.ToJSON()) // {"Tags":["足球","篮球","排球"],"Uid":133,"Uname":"An"}

	//var arr[]string = []string{"a", "b", "c"}
	//fmt.Println(arr)
	//fmt.Println(len(arr))

	//var arr [3]string
	//arr[0] = "a"
	//arr[2] = "c"
	//fmt.Println(arr)
	//fmt.Println(len(arr))

	//var arr [3]string = [3]string{0: "a", 2: "c"}
	//fmt.Println(arr)
	//fmt.Println(len(arr))
}
