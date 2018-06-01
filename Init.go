package xormTool

import (
	"github.com/xormplus/xorm"
)
//全局Db实例
var Db *xorm.Engine
//全局Session实例
var LocalSession *xorm.Session

//并不单纯指连接池，当需要链接多个数据库的多个db,比如Dbs["medium"]用的是medium数据库,Dbs["slot"]用的是slot数据库
var Dbs map[string]*xorm.Engine
//对应Dbs不同数据库的全局session
var LocalSessions map[string]*xorm.Session

func init(){
	Dbs = make(map[string]*xorm.Engine)
	LocalSessions = make(map[string]*xorm.Session)
}

//配置数据源，单源
func DataSource(dataSource string) (*xorm.Engine,error){
	var err error
	Db,err=xorm.NewPostgreSQL(dataSource)
	if err!=nil{
		return nil,err
	}
	LocalSession = Db.NewSession()

	Dbs["default"] = Db
	LocalSessions["default"] = LocalSession
	return Db,nil
}

//新增数据源,以键值保存到全局Dbs里
func NewDataSource(key string,dataSource string)(*xorm.Engine,error){
	var err error
	DbNew,err:=xorm.NewPostgreSQL(dataSource)
	if err!=nil{
		return nil,err
	}
	LocalSessionNew := Db.NewSession()


	Dbs[key] = DbNew
	LocalSessions[key] = LocalSessionNew
	return DbNew,nil
}

//实例属性配置，可设置是否显示sql语句，最大连接数和最大闲置数
func Config(printSQL bool,maxIdleConns int,maxOpenConns int){
	for k,_ :=range Dbs{
		Dbs[k].ShowSQL(printSQL)
		Dbs[k].SetMaxOpenConns(maxOpenConns)
		Dbs[k].SetMaxIdleConns(maxIdleConns)
	}
	////2.显示sql语句
	//Db.ShowSQL(printSQL)
	////3.设置连接数
	//Db.SetMaxIdleConns(maxIdleConns)
	//Db.SetMaxOpenConns(maxOpenConns)
}

//默认的db实例配置
func DefaultConfig(){
	for k,_ :=range Dbs{
		Dbs[k].ShowSQL(true)
		Dbs[k].SetMaxOpenConns(2000)
		Dbs[k].SetMaxIdleConns(1000)
	}

	////2.显示sql语句
	//Db.ShowSQL(true)
	//
	////3.设置连接数
	//Db.SetMaxIdleConns(2000)
	//Db.SetMaxOpenConns(1000)
}

//
func GetDb(key string) *xorm.Engine{
	if key==""{
		return Db
	}
	return Dbs[key]
}