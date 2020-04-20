package services

import "basePkg/models"

// 想要不同包调用 func 开头必须大写
func GetUser() string {
	//var news string = getNew() // 同包调用不是引入
	//return "return string" + news
	user := new(models.UserModel)
	user.Uname = "An"
	user.SetVal(12, "hello") // 指针问题
	return user.ToString()
}