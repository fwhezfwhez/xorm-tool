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

	result,err:=Select("select * from class")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(result)

}
