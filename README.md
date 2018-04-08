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
    	db.DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
        db.DefaultConfig()	
        //config specifically
        // db.Config(true,2000,1000)
        
    	//1.insert
    	id,err:=db.Insert("insert into class(name) values(?) returning id","测试数据")
    	if err!=nil{
    		fmt.Println(err)
    		return
    	}
    	fmt.Println(id)
        
    	//2.update
    	id,err=db.Update("update class set name=? where id=? returning id","测试数据",23)
    	if err!=nil{
    		fmt.Println(err)
    		return
    	}
    	fmt.Println(id)
        
    	//3.delete
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
    	
    	//4.动态查询条件DynamicSelect
    	    //4.1DynamicSelect
    		type User struct{
        		Id int
        		Name string
        		ClassId int `db:"class_id"`
        		Description string
        	}
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
        	err=db.DynamicSelect(&users,"select * from public.user",whereMap,orderBy,"desc",2,0,"ft4",1035,1)
        	if err!=nil{
        		fmt.Println(err)
        		return
        	}
        	fmt.Println(users)
    	    
        	//4.2DynamicSelectOne
        	    user :=User{}
            	whereMap2 := make([][]string,0)
            	whereMap2 = append(whereMap2,[]string{
            		"","id","=",
            	})
            	err=db.DynamicSelectOne(&user,"select * from public.user",whereMap2,nil,"desc",2,0,1036)
            	if err!=nil{
            		fmt.Println(err)
            		return
            	}
            	fmt.Println(user)
            
            //4.3 DynamicSelectCount
            	orderBy3 :=make([]string,0)
            	orderBy3 = append(orderBy3,"id")
            	whereMap3 := make([][]string,0)
            	whereMap3 = append(whereMap3,[]string{
            		"","name","=",
            	})
            	whereMap3 = append(whereMap3,[]string{
            		"and","id",">",
            	})
            	whereMap3 = append(whereMap3,[]string{
            		"and","id","<>",
            	})
            	count,err:=db.DynamicSelectCount("select * from public.user",whereMap3,nil,"desc",2,0,"ft4",1035,1)
            	if err!=nil{
            		fmt.Println(err)
            		return 
            	}
            	fmt.Println(count)
    	//个性化操作
    	db:=db.GetDb()
    	s:=db.NewSession()
    	...
    	//do sth specail
    }
```