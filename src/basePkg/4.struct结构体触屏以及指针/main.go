package main

import "fmt"

type NewsModel struct {
	NewsId int
	NewsTitle, NewsContent string
}
// new 开辟内存、初始化
func main() {
	//var news NewsModel
	//news.NewsId = 1
	//news.NewsTitle = "title"
	//news.NewsContent = "content"
	//fmt.Println(news)

	//var news NewsModel = NewsModel{123, "title", "content"}
	//fmt.Println(news)

	//var news NewsModel = NewsModel{NewsContent:"content", NewsId:123, NewsTitle:"aaa"}
	//fmt.Println(news)

	var news *NewsModel
	news=new(NewsModel)
	//*news = NewsModel{123, "aa", "aa"}
	//news = &NewsModel{123, "aa", "aa"}
	news.NewsTitle = "title"
	(*news).NewsId = 123
	fmt.Println(news)
	fmt.Println(news.NewsContent)
	fmt.Println(&(news.NewsContent))
}
