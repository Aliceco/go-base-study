package models

type UserModel struct {
	Uid int
	Uname string
}
// 纯方法
func ToString() string {
	return "测试字符串"
}
// 这个是类方法
func (u UserModel) ToString() string  {
	return "用户名是" + u.Uname
}

func (u UserModel) SetVal(id int, name string) {
	u.Uname = name
	u.Uid = id
}