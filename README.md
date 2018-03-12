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
    
    	//查询,查询解耦不出！！！！提不出api,直接拿原生session操作吧
    	type Class struct {
    		Id int
    		Name string
    	}
    	session :=db.Db.NewSession()
    	class := make([]Class,0)
    	session.Begin()
    	err=session.SQL("select * from class").Find(&class)
    	if err!=nil {
    		session.Rollback()
    		panic(err)
    	}
    	fmt.Println(class)
    	session.Commit()
    }
```