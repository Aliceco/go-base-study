package main

import "basePkg/myHttpServer/core"

type NewsController struct {
	core.MyController
}

func(this *NewsController) GET(){
	this.Context.WriterString("this in newsController GET")
}

func(this *NewsController) POST(){
	this.Context.WriterString("this in newsController POST")
}