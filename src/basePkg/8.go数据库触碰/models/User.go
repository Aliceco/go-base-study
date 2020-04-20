package models

type UserModel struct {
	Uname string
	Upassword string
}
// 纯方法
func ToString() string {
	return "测试字符串"
}
// 这个是类方法
func (u UserModel) ToString() string  {
	return "用户名是" + u.Uname
}