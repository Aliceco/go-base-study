package core

import (
	"github.com/pquerna/ffjson/ffjson"
	"io/ioutil"
)

type MyController struct {
	Context *MyContext
}

func (this *MyController) Init(Context *MyContext) {
	this.Context = Context
}
// 获取get参数代码
func (this *MyController) GetParam(key string, defValue ...string) string {
	res := this.Context.request.URL.Query().Get(key)
	if res==""&&len(defValue)>0 {
		return defValue[0]
	}
	return res
}
// 获取post参数代码
func (this *MyController) PostParam(key string, defValue ...string) string {
	res := this.Context.request.PostFormValue(key) // 单参数
	if res==""&&len(defValue)>0 {
		return defValue[0]
	}
	return res
}
// 获取JSON参数代码
func (this *MyController) JsonParam(obj interface{}) {
	body, err:=ioutil.ReadAll(this.Context.request.Body)
	if err!=nil {
		panic(err)
	}
	err=ffjson.Unmarshal(body, obj)
	if err!=nil {
		panic(err)
	}
}

type ControllerInterface interface {
	Init(Context *MyContext)
	GET()
	POST()
}