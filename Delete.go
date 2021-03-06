package xormTool
//删除数据并返回删除的id
//Delete("delete from tUser where id=?",5)   returns 5,nil | -1 ,err
func Delete(sql string,args...interface{})(int,error){
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
