package xormTool

import (
	"testing"

	"strconv"
)

func TestInsert(t *testing.T) {
	sql := "insert into class(name) values(?) "
	id,err:=Insert(sql,"高中6班")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(id)
}
func TestLocalSessionInsert(t *testing.T) {
	DataSource("postgres://medium:mediuml4eLxglxL8@111.231.137.127:5432/fe?sslmode=disable")
	Db.ShowSQL(true)
	session:=Db.NewSession()
	var id int
	var temp string
	var err error
	for i:=0;i<1000;i++{
		temp=strconv.Itoa(i)
		_,err=session.SQL("insert into medium_slot(slotname,slotid) values(?,?) returning id",temp,temp).Get(&id)
		if err!=nil{
			t.Log(err)
		}
		t.Log(id)
	}

	session.Commit()
	//sql := "insert into medium_medium(mediumname) values(?) returning id"
	//id,err:=LocalSessionInsert(sql,"高中6班")

}
