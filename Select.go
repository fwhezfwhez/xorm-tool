package xormTool

import (
	"strconv"
	"strings"
	"fmt"
	"errors"
)

type Class struct {
	Id   int
	Name string
}

func Select(dest interface{}, sql string, args ...interface{}) (error) {
	session := Db.NewSession()
	session.Begin()
	//*[]xormToll.Class
	//怎么转化为*[]xormToll.Class
	err := session.SQL(sql, args...).Find(dest)
	if err != nil {
		session.Rollback()
		return err
	}
	session.Commit()
	return nil
}

func SelectByDb(key string, dest interface{},sql string,args ...interface{})(error){
	if dbTemp,ok :=Dbs[key];!ok{
		return errors.New("没有找到key为"+key+"的数据源，使用db.NewDataSource(key,'xxxxx')来添加多个数据源")
	}else {
		session := dbTemp.NewSession()
		session.Begin()
		err := session.SQL(sql, args...).Find(dest)
		if err != nil {
			session.Rollback()
			return err
		}
		session.Commit()
		return nil
	}
}
func SelectOne(dest interface{}, sql string, args ...interface{}) error {
	session := Db.NewSession()
	defer session.Close()
	session.Begin()
	_, err := session.SQL(sql, args...).Get(dest)
	if err != nil {
		session.Rollback()
		return err
	}
	session.Commit()
	return nil
}

func SelectOneByDb(key string, dest interface{},sql string,args ...interface{})(error){
	if dbTemp,ok :=Dbs[key];!ok{
		return errors.New("没有找到key为"+key+"的数据源，使用db.NewDataSource(key,'xxxxx')来添加多个数据源")
	}else {
		session := dbTemp.NewSession()
		session.Begin()
		_,err := session.SQL(sql, args...).Get(dest)
		if err != nil {
			session.Rollback()
			return err
		}
		session.Commit()
		return nil
	}
}

//
func SelectCount(sql string, args ...interface{}) (int, error) {
	session := Db.NewSession()
	defer session.Close()
	session.Begin()
	count, err := session.SQL(sql, args...).Query().Count()
	if err != nil {
		session.Rollback()
		return -1, err
	}
	session.Commit()
	return count, nil
}

func DynamicSelectByDb(key string,dest interface{}, basicSql string, whereMap [][]string, orderBy []string, Asc string, limit int, offset int, args ...interface{}) (error) {
	sql := rollingSql(basicSql, whereMap, orderBy, Asc, limit, offset)
	args = removeZero(args)
	return SelectByDb(key,dest, sql, args...)
}
func DynamicSelect(dest interface{}, basicSql string, whereMap [][]string, orderBy []string, Asc string, limit int, offset int, args ...interface{}) (error) {
	sql := rollingSql(basicSql, whereMap, orderBy, Asc, limit, offset)
	args = removeZero(args)
	return Select(dest, sql, args...)
}

func DynamicSelectOneByDb(key string,dest interface{}, basicSql string, whereMap [][]string, orderBy []string, Asc string, limit int, offset int, args ...interface{}) (error) {
	sql := rollingSql(basicSql, whereMap, orderBy, Asc, limit, offset)
	args = removeZero(args)
	return SelectOne(dest, sql, args...)
}
func DynamicSelectOne(dest interface{}, basicSql string, whereMap [][]string, orderBy []string, Asc string, limit int, offset int, args ...interface{}) (error) {
	sql := rollingSql(basicSql, whereMap, orderBy, Asc, limit, offset)
	args = removeZero(args)
	return SelectOne(dest, sql, args...)
}

func DynamicJoinSelectByDb(key string,dest interface{},basicSql string,whereMap [][]string,joinMap []string,orderBy []string,asc string, limit int,offset int,args ...interface{})(error){
	//basicSql="select a.* from payorder a,productid_ads b "
	sql :=basicSql
	//连接查询条件，注意joinMap的值是否有空格和and由输入自行控制，比如joinMap=[" a.id=b.userId ","and a.id=5"]
	if joinMap!=nil ||len((joinMap))!=0{
		for _,v:=range joinMap{
			sql = sql + " where "+v
		}
	}

	if len(whereMap) != 0 {
		if !(joinMap!=nil ||len((joinMap))!=0) {
			sql = sql + " where "
		}
		for _, v := range whereMap {
			//v[0]表示性质，and 还是or,v[1]表示field，比如name，age,v[2]表示条件符号,=,>,<,<>,like
			if v[2] == "between" {
				sql = sql + " " + v[0] + " " + v[1] + " " + "between" + " " + "?" + " " + "and" + " " + "?" + " "
				continue
			}

			sql = sql + " " + v[0] + " " + v[1] + " " + v[2] + " " + "?"
		}
	}
	//fmt.Println("处理where完毕:"+sql)

	//2.处理Orderby和asc
	if len(orderBy) != 0 && orderBy != nil {
		sql = sql + " order by " + strings.Join(orderBy, ",") + " " + asc + " "
	}
	//fmt.Println("处理order,asc完毕:"+sql)

	//3.处理limit,offset
	if limit != -1 && offset != -1 {
		sql = sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)
	}
	fmt.Println(sql)
	args = removeZero(args)
	return SelectByDb(key,dest, sql, args...)
}
func DynamicJoinSelect(dest interface{},basicSql string,whereMap [][]string,joinMap []string,orderBy []string,asc string, limit int,offset int,args ...interface{})(error){
	//basicSql="select a.* from payorder a,productid_ads b "
	sql :=basicSql
	//连接查询条件，注意joinMap的值是否有空格和and由输入自行控制，比如joinMap=[" a.id=b.userId ","and a.id=5"]
		if joinMap!=nil ||len((joinMap))!=0{
			for _,v:=range joinMap{
				sql = sql + " where "+v
			}
		}

		if len(whereMap) != 0 {
			if !(joinMap!=nil ||len((joinMap))!=0) {
				sql = sql + " where "
			}
			for _, v := range whereMap {
				//v[0]表示性质，and 还是or,v[1]表示field，比如name，age,v[2]表示条件符号,=,>,<,<>,like
				if v[2] == "between" {
					sql = sql + " " + v[0] + " " + v[1] + " " + "between" + " " + "?" + " " + "and" + " " + "?" + " "
					continue
				}

				sql = sql + " " + v[0] + " " + v[1] + " " + v[2] + " " + "?"
			}
		}
		//fmt.Println("处理where完毕:"+sql)

		//2.处理Orderby和asc
		if len(orderBy) != 0 && orderBy != nil {
			sql = sql + " order by " + strings.Join(orderBy, ",") + " " + asc + " "
		}
		//fmt.Println("处理order,asc完毕:"+sql)

		//3.处理limit,offset
		if limit != -1 && offset != -1 {
			sql = sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)
		}
		fmt.Println(sql)
	    args = removeZero(args)
	    return Select(dest, sql, args...)
}

func DynamicJoinSelectOneByDb(key string,dest interface{},basicSql string,whereMap [][]string,joinMap []string,orderBy []string,asc string, limit int,offset int,args ...interface{})(error){
	//basicSql="select a.* from payorder a,productid_ads b "
	sql :=basicSql
	//连接查询条件，注意joinMap的值是否有空格和and由输入自行控制，比如joinMap=[" a.id=b.userId ","and a.id=5"]
	if joinMap!=nil ||len((joinMap))!=0{
		for _,v:=range joinMap{
			sql = sql + " where "+v
		}
	}

	if len(whereMap) != 0 {
		if !(joinMap!=nil ||len((joinMap))!=0) {
			sql = sql + " where "
		}
		for _, v := range whereMap {
			//v[0]表示性质，and 还是or,v[1]表示field，比如name，age,v[2]表示条件符号,=,>,<,<>,like
			if v[2] == "between" {
				sql = sql + " " + v[0] + " " + v[1] + " " + "between" + " " + "?" + " " + "and" + " " + "?" + " "
				continue
			}

			sql = sql + " " + v[0] + " " + v[1] + " " + v[2] + " " + "?"
		}
	}
	//fmt.Println("处理where完毕:"+sql)

	//2.处理Orderby和asc
	if len(orderBy) != 0 && orderBy != nil {
		sql = sql + " order by " + strings.Join(orderBy, ",") + " " + asc + " "
	}
	//fmt.Println("处理order,asc完毕:"+sql)

	//3.处理limit,offset
	if limit != -1 && offset != -1 {
		sql = sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)
	}
	fmt.Println(sql)
	args = removeZero(args)
	return SelectOneByDb(key,dest, sql, args...)
}


func DynamicJoinSelectOne(dest interface{},basicSql string,whereMap [][]string,joinMap []string,orderBy []string,asc string, limit int,offset int,args ...interface{})(error){
	//basicSql="select a.* from payorder a,productid_ads b "
	sql :=basicSql
	//连接查询条件，注意joinMap的值是否有空格和and由输入自行控制，比如joinMap=[" a.id=b.userId ","and a.id=5"]
	if joinMap!=nil ||len((joinMap))!=0{
		for _,v:=range joinMap{
			sql = sql + " where "+v
		}
	}

	if len(whereMap) != 0 {
		if !(joinMap!=nil ||len((joinMap))!=0) {
			sql = sql + " where "
		}
		for _, v := range whereMap {
			//v[0]表示性质，and 还是or,v[1]表示field，比如name，age,v[2]表示条件符号,=,>,<,<>,like
			if v[2] == "between" {
				sql = sql + " " + v[0] + " " + v[1] + " " + "between" + " " + "?" + " " + "and" + " " + "?" + " "
				continue
			}

			sql = sql + " " + v[0] + " " + v[1] + " " + v[2] + " " + "?"
		}
	}
	//fmt.Println("处理where完毕:"+sql)

	//2.处理Orderby和asc
	if len(orderBy) != 0 && orderBy != nil {
		sql = sql + " order by " + strings.Join(orderBy, ",") + " " + asc + " "
	}
	//fmt.Println("处理order,asc完毕:"+sql)

	//3.处理limit,offset
	if limit != -1 && offset != -1 {
		sql = sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)
	}
	fmt.Println(sql)
	args = removeZero(args)
	return SelectOne(dest, sql, args...)
}

func DynamicSelectCount(basicSql string, whereMap [][]string, orderBy []string, Asc string, limit int, offset int, args ...interface{}) (int, error) {
	sql := rollingSql(basicSql, whereMap, orderBy, Asc, limit, offset)
	args = removeZero(args)
	return SelectCount(sql, args...)
}

func rollingSql(basicSql string, whereMap [][]string, orderBy []string, Asc string, limit int, offset int) string {
	var sql = basicSql
	//1.处理where
	if len(whereMap) != 0 {
		sql = sql + " where "
		for _, v := range whereMap {
			//v[0]表示性质，and 还是or,v[1]表示field，比如name，age,v[2]表示条件符号,=,>,<,<>,like
			if v[2] == "between" {
				sql = sql + " " + v[0] + " " + v[1] + " " + "between" + " " + "?" + " " + "and" + " " + "?" + " "
				continue
			}

			sql = sql + " " + v[0] + " " + v[1] + " " + v[2] + " " + "?"
		}
	}
	//fmt.Println("处理where完毕:"+sql)

	//2.处理Orderby和asc
	if len(orderBy) != 0 && orderBy != nil {
		sql = sql + " order by " + strings.Join(orderBy, ",") + " " + Asc + " "
	}
	//fmt.Println("处理order,asc完毕:"+sql)

	//3.处理limit,offset
	if limit != -1 && offset != -1 {
		sql = sql + " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)
	}
	//fmt.Println(sql)
	return sql
}

func remove(slice []interface{}, elem interface{}) []interface{}{
	if len(slice) == 0 {
		return slice
	}
	for i, v := range slice {
		if v == elem {
			slice = append(slice[:i], slice[i+1:]...)
			return remove(slice,elem)
			break
		}
	}
	return slice
}
func removeZero(slice []interface{}) []interface{}{
	if len(slice) == 0 {
		return slice
	}
	for i, v := range slice {
		if ifZero(v) {
			slice = append(slice[:i], slice[i+1:]...)
			return removeZero(slice)
			break
		}
	}
	return slice
}

func ifZero(arg interface{}) bool {
	if arg==nil{
		return true
	}
	switch v := arg.(type) {
	case int, float64, int32, int16, int64, float32:
		if v == 0 {
			return true
		}
	case string:
		if v == "" || v == "%%" ||v=="%"{
			return true
		}
	default:
		return false
	}
	return false
}
