package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
//type all interface {}
//var all interface{} = "123"
func main() {
	// 空接口
	//var i interface{} = "123"
	//fmt.Println(i)

	// 空切片
	//list:=make([]interface{}, 2)
	//list[0] = 1
	//list[1] = "abc"
	//fmt.Println(list)

	//map使用
	//m:=map[string]int{"userName": 123}
	//fmt.Println(m)

	//m:=map[string]string{"userName": "An"}
	//fmt.Println(m)

	m:=make(map[string]int)
	m["userName"] = 12
	m["userAge"] = 12
	delete(m, "userName")
	fmt.Println(m)


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
	//userModel := []models.UserModel{}
	allRow:=make([]interface{}, 0)

	for rows.Next() {
		var uname,upassword string
		rows.Scan(&uname, &upassword)
		//oneRow:=make([]interface{}, 2)
		//oneRow[0]=uname
		//oneRow[1]=upassword
		//allRow = append(allRow, oneRow)
		allRow = append(allRow, []interface{}{uname, upassword})
	}
	fmt.Println(allRow)
}
