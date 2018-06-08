package xormTool

import (
	_ "github.com/lib/pq"
	"errors"
)

/*
	1.插入sql语句和参数，insert返回新增行的id和错误信息
 */
func Insert(sql string,args ...interface{}) (int,error){
	session := Db.NewSession()
	defer session.Close()
	session.Begin()
	var  id int
	_,err:=session.SQL(sql,args...).Get(&id)
	if err!=nil {
		session.Rollback()
		return -1,err
	}
	session.Commit()

	return id,err
}

//选择db对象来执行插入操作
func InsertByDb(key string,sql string,args ...interface{})(int,error){
	if dbTemp,ok :=Dbs[key];!ok{
		return -1,errors.New("没有找到key为"+key+"的数据源，使用db.NewDataSource(key,'xxxxx')来添加多个数据源")
	}else{
		session:=dbTemp.NewSession()
		defer session.Close()
		session.Begin()
		var  id int
		_,err:=session.SQL(sql,args...).Get(&id)
		if err!=nil {
			session.Rollback()
			return -1,err
		}
		session.Commit()

		return id,err
	}

}

func LocalSessionInsert(sql string,args ...interface{}) (int,error){
	LocalSession.Begin()
	var  id int
	_,err:=LocalSession.SQL(sql,args...).Get(&id)
	if err!=nil {
		LocalSession.Rollback()
		return -1,err
	}
	LocalSession.Commit()
	return id,nil
}