package xormTool

import (
	"testing"
)

func TestUpdate(t *testing.T) {
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	id,err:=Update("update class set name=? where id=? returning id","高中n班",1)
	if err!=nil {
		t.Fatal(err)
	}
	t.Log(id)
}

//
func TestDynamicUpdate(t *testing.T) {
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	Db.ShowSQL(true)
	type U struct{
		Name string `xorm:"name"`
		Id  int `xorm:"id"`
	}
	u:=U{"高中一班",2}
	whereMap:=make([][]string,0)
	whereMap = append(whereMap,[]string{
		"","id","=",
	})
	id,err:=DynamicUpdate("default","update class",u,whereMap,"高中m班",9999,1)
	if err!=nil{
	t.Fatal(err)
	}
	t.Log(id)
}