package main

import (
	"basePkg/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 数组切片学习
	//arr := [5]int {1, 2, 3}
	//fmt.Println(arr)
	//fmt.Println(len(arr)) // len数组长度
	//fmt.Println(cap(arr)) // cap数组容量

	//arr := []int {1, 2, 3}
	//fmt.Println(arr)
	//fmt.Println(len(arr)) // len数组长度
	//fmt.Println(cap(arr))

	// 有固定长度就是array、没有固定长度就是切片
	//arr := [5]int{1,2,3,4,5}
	//arr := make([]int, 5) // make分配空间、初始化长度
	//fmt.Println(arr) // [0 0 0 0 0]
	//fmt.Println(len(arr)) // 5
	//fmt.Println(cap(arr)) // 5

	//arr := [5]int{1,2,3,4,5}
	//// 左闭( [] )右开( [) )     [4, 10)  4<10
	//s:=arr[0:4]
	//s = append(s, 3, 4, 5, 6) // 如果cap不够则会创建一个新的slice，把原有的slice拷贝进去（实际开发要预估cap，否则会导致性能开销）
	//fmt.Println(s) // [1 2 3 4]
	//fmt.Println(len(s))
	//fmt.Println(cap(s))

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
	userModel := []models.UserModel{}
	for rows.Next() {
		tmp := models.UserModel{}
		rows.Scan(&tmp.Uname, &tmp.Upassword)
		userModel = append(userModel, tmp)
	}
	fmt.Println(userModel)
}
