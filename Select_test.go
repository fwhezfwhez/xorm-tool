package xormTool

import (
	"testing"
)

func TestSelectCount(t *testing.T) {
	count,err:=SelectCount("select * from class")
	if err!=nil {
		t.Fatal(err)
	}
	t.Log(count)
}


func TestSelect(t *testing.T) {
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
