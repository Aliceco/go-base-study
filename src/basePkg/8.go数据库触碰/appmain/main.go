package main

import (
	"basePkg/models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(192.168.99.100:3306)/go_base?charset=utf8")
	if err!=nil {
		fmt.Println("错误连接"+err.Error())
		return
	}
	rows,err:=db.Query("select name,password from user_test")
	if err!=nil {
		fmt.Println("错误查询"+err.Error())
		return
	}
	fmt.Println(rows.Columns())

	//var userModel models.UserModel = models.UserModel{}
	userModel := models.UserModel{}
	for rows.Next() {
		//var upassword string
		//var uname string
		//rows.Scan(&uname, &upassword)
		//fmt.Println(uname, upassword)
		rows.Scan(&userModel.Uname, &userModel.Upassword)
		fmt.Println(userModel)
	}
}
