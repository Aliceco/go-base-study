package main

import (
	"basePkg/1.go包的触碰/models"
	"basePkg/myHttpServer/core"
)

type NewsController struct {
	core.MyController
}

func(this *NewsController) GET(){
	//obj:=map[string]string{}
	//obj["userName"] = "an"
	//obj:=core.MyMap{
	//	"user": "an",
	//	"age": 12,
	//}
	//this.Context.WriteJSON(obj)

	//this.Context.WriterString("this in newsController GET")

	p:=this.GetParam("user", "no param")
	// http://localhost:8099/?user=1111
	this.Context.WriterString(p)
}

func(this *NewsController) POST(){
	//this.Context.WriterString("this in newsController POST")

	//p:=this.PostParam("user", "no param")
	//this.Context.WriterString(p)

	user:=models.UserModel{}
	this.JsonParam(&user)
	this.Context.WriteJSON(user)
}