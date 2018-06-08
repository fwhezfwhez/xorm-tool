package xormTool

import (
	"errors"
)

func Update(sql string,args ...interface{})(int,error){
	var id int

	session :=Db.NewSession()
	session.Begin()
	defer session.Close()
	_,err:=session.SQL(sql,args...).Get(&id)
	if err!=nil {
		session.Rollback()
		return -1,err
	}
	session.Commit()
	return id,nil
}

func UpdateByDb(key string ,sql string, args ...interface{})(int,error){
	if dbTemp,ok :=Dbs[key];!ok{
		return -1,errors.New("没有找到key为"+key+"的数据源，使用db.NewDataSource(key,'xxxxx')来添加多个数据源")
	}else {

		var id int
		session:=dbTemp.NewSession()
		session.Begin()
		defer session.Close()
		_,err:=session.SQL(sql,args...).Get(&id)
		if err!=nil {
			session.Rollback()
			return -1,err
		}
		session.Commit()
		return id,nil
	}
}