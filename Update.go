package xormTool

import (
	"errors"
	"strings"
	"reflect"
)

func Update(sql string, args ...interface{}) (int, error) {
	var id int

	session := Db.NewSession()
	session.Begin()
	defer session.Close()
	_, err := session.SQL(sql, args...).Get(&id)
	if err != nil {
		session.Rollback()
		return -1, err
	}
	session.Commit()
	return id, nil
}

func UpdateByDb(key string, sql string, args ...interface{}) (int, error) {
	if dbTemp, ok := Dbs[key]; !ok {
		return -1, errors.New("没有找到key为" + key + "的数据源，使用db.NewDataSource(key,'xxxxx')来添加多个数据源")
	} else {

		var id int
		session := dbTemp.NewSession()
		session.Begin()
		defer session.Close()
		_, err := session.SQL(sql, args...).Get(&id)
		if err != nil {
			session.Rollback()
			return -1, err
		}
		session.Commit()
		return id, nil
	}
}

//unsupport for time.Time arg as set xxx=xxx
func DynamicUpdate(Dbkey string, sql string, updateBasic interface{}, whereMap [][]string, args ... interface{}) (int, error) {

	if dbTemp, ok := Dbs[Dbkey]; !ok {
		return -1, errors.New("没有找到key为" + Dbkey + "的数据源，使用db.NewDataSource(key,'xxxxx')来添加多个数据源,使用Dbkey='default'获得DataSource()API默认数据源",)
	} else {
		//update xx  set xx=xx,xx=xx,xx=xx... where xx=xx
		if !(strings.Contains(sql, "set") || strings.Contains(sql, "SET")) {
			sql = sql + " set "
		}
		var tagValueTemp string
		var valueTmp interface{}
		//var nameTmp string
		var rsTmp = true
		var paramTmp = make([]string, 0)
		v := reflect.ValueOf(updateBasic)
		//t := reflect.TypeOf(updateBasic)
		for i := 0; i < v.NumField(); i++ {
			tagValueTemp = v.Type().Field(i).Tag.Get("xorm")

			if tagValueTemp == "" {
				continue
			}
			valueFieldTmp := v.Field(i)
			//nameTmp = t.Field(i).Name
			switch valueFieldTmp.Type().String() {
			case "string":
				valueTmp = valueFieldTmp.String()
				if valueTmp == "" {
					rsTmp = false
				}
			case "int", "int8", "int16", "int64", "int32":
				valueTmp = valueFieldTmp.Int()
				if valueTmp == 0 {
					rsTmp = false
				}
			case "float", "float32", "float64":
				valueTmp = valueFieldTmp.Float()
				if valueTmp == float64(0) {
					rsTmp = false
				}
			case "*string", "*int", "*int64", "*int32", "*int16", "*int8", "*float", "*float32", "*float64":
				if valueFieldTmp.IsNil() {
					rsTmp = false
				}
			}
			if rsTmp {
				paramTmp = append(paramTmp, tagValueTemp+"=?")
			}
		}
		sql = sql + strings.Join(paramTmp, " , ")
		sql = RollingSql(sql, whereMap, nil, "", -1, -1)

		args = RemoveZero(args)
		var id int

		_, er := dbTemp.SQL(sql+"returning id", args...).Get(&id)
		if er != nil {
			return -1, er
		}
		return id, nil
	}
}
