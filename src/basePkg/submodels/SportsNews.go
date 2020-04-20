package submodels

import (
	"basePkg/models"
	"github.com/pquerna/ffjson/ffjson"
)

type SportsNew struct {
	Tags []string // 字符串数组
	//News models.UserModel
	models.UserModel
}
func (sn SportsNew) ToJSON() string {
	res, err := ffjson.Marshal(sn)
	if err != nil {
		return err.Error()
	} else {
		// 返回的[]byte所以转成string
		return string(res)
	}
}