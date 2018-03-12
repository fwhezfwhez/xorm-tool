package xormTool

import (
)

type Class struct{
	Id int
	Name string
}

/*
		Class无法解耦！！
 */
func Select(sql string,args...interface{}) ([]Class,error){
	session:=Db.NewSession()
	session.Begin()
	//*[]xormToll.Class
	dest := make([]Class,0)

	//怎么转化为*[]xormToll.Class
	err:=session.SQL(sql,args...).Find(&dest)
	if err!=nil{
		session.Rollback()
		return nil,err
	}
	session.Commit()
	return dest,nil
}

func SelectCount(sql string,args...interface{})(int,error){
	session:=Db.NewSession()
	defer session.Close()
	session.Begin()
	count,err :=session.SQL(sql,args...).Query().Count()
	if err!=nil {
		session.Rollback()
		return -1,err
	}
	session.Commit()
	return count,nil
}