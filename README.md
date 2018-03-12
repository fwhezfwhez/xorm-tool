a basic pattern to use xorm package
##引用包:
"github.com/xormplus/xorm"

##使用方式
go get github.com/fwhezfwhez/xorm-tool

##Example
```go

    sql := "insert into class(name) values(?) "
    	id,err:=Insert(sql,"高中6班")
    	if err!=nil{
    		t.Fatal(err)
    	}
    	t.Log(id)
```