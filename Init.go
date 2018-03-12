package xormTool

import (
	"github.com/xormplus/xorm"
	"fmt"
)
var Db *xorm.Engine
var LocalSession *xorm.Session

func init(){
	var err error
	Db, err = xorm.NewPostgreSQL("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	//2.显示sql语句
	Db.ShowSQL(true)

	//3.设置连接数
	Db.SetMaxIdleConns(2000)
	Db.SetMaxOpenConns(1000)
	//type Id struct{
	//	id int
	//}
	//var newId = Id{}
	Db.Ping()
	LocalSession = Db.NewSession()
}
