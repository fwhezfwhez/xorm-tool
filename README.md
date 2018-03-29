a basic pattern to use xorm package
##引用包:
"github.com/xormplus/xorm"

##使用方式
go get github.com/fwhezfwhez/xorm-tool

##Example
```go

    package main
    import(
    	db "github.com/fwhezfwhez/xorm-tool"
    	"fmt"
    )
    func main(){
    
    	//insert
    	db.DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
    	db.DefaultConfig()
    	id,err:=db.Insert("insert into class(name) values(?) returning id","测试数据")
    	if err!=nil{
    		fmt.Println(err)
    		return
    	}
    	fmt.Println(id)
    
    	id,err=db.Update("update class set name=? where id=? returning id","测试数据",23)
    	if err!=nil{
    		fmt.Println(err)
    		return
    	}
    	fmt.Println(id)
    
    	id,err=db.Delete("update class set name=? where id=? returning id","测试数据",23)
    	if err!=nil{
    		fmt.Println(err)
    		return
    	}
    	fmt.Println(id)
    	
    	type Class struct {
    		Id int
    		Name string
    	}
    	
    	classes := make([]Class,0)
    	db.Select(&classes,"select * from class where id>?",2)
    	fmt.Println(classes)
    	
    	class :=Class{}
    	db.SelectOne(&class,"select * from class where id=?",2)
    	fmt.Println(class)
    	
    	//个性化操作
    	db:=db.GetDb()
    	s:=db.NewSession()
    	...
    	//do sth specail
    }
```