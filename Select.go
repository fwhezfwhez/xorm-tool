package xormTool

import (
	"strconv"
	"fmt"
	"strings"
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

func DynamicSelect(dest interface{},basicSql string,whereMap [][]string,orderBy []string,Asc string,limit int,offset int,args...interface{})(error){
	sql :=rollingSql(basicSql,whereMap,orderBy,Asc,limit,offset,args...)
	return Select(dest,sql,args...)
}

func DynamicSelectOne(dest interface{},basicSql string,whereMap [][]string,orderBy []string,Asc string,limit int,offset int,args...interface{})(error){
	sql :=rollingSql(basicSql,whereMap,orderBy,Asc,limit,offset,args...)
	return SelectOne(dest,sql,args...)
}

func DynamicSelectCount(basicSql string,whereMap [][]string,orderBy []string,Asc string,limit int,offset int,args...interface{})(int,error){
	sql :=rollingSql(basicSql,whereMap,orderBy,Asc,limit,offset,args...)
	return SelectCount(sql,args...)
}

func rollingSql(basicSql string,whereMap [][]string,orderBy []string,Asc string,limit int,offset int,args...interface{})string{
	var sql =basicSql
	//1.处理where
	if len(whereMap)!=0 {
		sql =sql+" where "
		for _,v:=range whereMap{
			//v[0]表示性质，and 还是or,v[1]表示field，比如name，age,v[2]表示条件符号,=,>,<,<>,like
			sql = sql +" "+ v[0]+" "+v[1]+v[2]+"?"
		}
	}
	//fmt.Println("处理where完毕:"+sql)

	//2.处理Orderby
	if len(orderBy)!=0 && orderBy!=nil{
		sql = sql+" order by "+strings.Join(orderBy,",")+" "+Asc+" "
	}
	//fmt.Println("处理order,asc完毕:"+sql)

	//3.处理limit,offset
	if limit!=-1&&offset!=-1{
		sql = sql + " limit "+strconv.Itoa(limit)+" offset "+strconv.Itoa(offset)
	}
	fmt.Println(sql)
	return sql
}