package xormTool

import (
	"testing"
)

func TestSelectCount(t *testing.T) {
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	count,err:=SelectCount("select * from class")
	if err!=nil {
		t.Fatal(err)
	}
	t.Log(count)
}


func TestSelect(t *testing.T) {
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	type Class struct {
		Id int
		Name string
	}
	classes :=make([]Class,0)
	err:=Select(&classes,"select * from class")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(classes)

}
func TestSelectOne(t *testing.T){
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	type Class struct {
		Id int
		Name string
	}
	class :=Class{}
	err:=SelectOne(&class,"select * from class where id=?",5)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(class)
}

func TestDynamicSelect(t *testing.T) {
	type User struct{
		Id int
		Name string
		ClassId int `db:"class_id"`
		Description string
	}
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	users :=make([]User,0)
	orderBy :=make([]string,0)
	orderBy = append(orderBy,"id")
	whereMap := make([][]string,0)
	whereMap = append(whereMap,[]string{
		"","name","=",
	})
	whereMap = append(whereMap,[]string{
		"and","id",">",
	})
	whereMap = append(whereMap,[]string{
		"and","id","<>",
	})
	err:=DynamicSelect(&users,"select * from public.user",whereMap,orderBy,"desc",2,0,"ft4",1035,1)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(users)
}

func TestDynamicSelectOne(t *testing.T) {
	type User struct{
		Id int
		Name string
		ClassId int `db:"class_id"`
		Description string
	}
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	user :=User{}
	orderBy :=make([]string,0)
	orderBy = append(orderBy,"id")
	whereMap := make([][]string,0)
	whereMap = append(whereMap,[]string{
		"","id","=",
	})
	err:=DynamicSelectOne(&user,"select * from public.user",whereMap,nil,"desc",2,0,1036)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(user)
}

func TestDynamicSelectCount(t *testing.T) {
	type User struct{
		Id int
		Name string
		ClassId int `db:"class_id"`
		Description string
	}
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	orderBy :=make([]string,0)
	orderBy = append(orderBy,"id")
	whereMap := make([][]string,0)
	whereMap = append(whereMap,[]string{
		"","name","=",
	})
	whereMap = append(whereMap,[]string{
		"and","id",">",
	})
	whereMap = append(whereMap,[]string{
		"and","id","<>",
	})
	count,err:=DynamicSelectCount("select * from public.user",whereMap,nil,"desc",2,0,"ft4",1035,1)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(count)
}