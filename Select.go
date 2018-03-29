package xormTool

import (
)

type Class struct{
	Id int
	Name string
}

func Select(dest interface{},sql string,args...interface{}) (error){
	session:=Db.NewSession()
	session.Begin()
	//*[]xormToll.Class
	//怎么转化为*[]xormToll.Class
	err:=session.SQL(sql,args...).Find(dest)
	if err!=nil{
		session.Rollback()
		return err
	}
	session.Commit()
	return nil
}

func SelectOne(dest interface{},sql string,args...interface{})error{
	session:=Db.NewSession()
	defer session.Close()
	session.Begin()
	_,err :=session.SQL(sql,args...).Get(dest)
	if err!=nil {
		session.Rollback()
		return err
	}
	session.Commit()
	return nil
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