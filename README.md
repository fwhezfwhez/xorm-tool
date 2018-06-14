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
			    var users=make([]User,0)
                arg1 :="ft"
				arg2 :=1035
				arg3 :=1
            	orderBy3 :=make([]string,0)
            	orderBy3 = append(orderBy3,"id")
            	whereMap3 := make([][]string,0)
				
				if arg1!=""{
				    whereMap3 = append(whereMap3,[]string{
            		"","name","=",
            	    })
				}

				if arg2!=0{
				    whereMap3 = append(whereMap3,[]string{
            		"","age","=",
            	    })
				}
				if arg3!=0{
				    whereMap3 = append(whereMap3,[]string{
            		"","class","=",
            	    })
				}				
            	count,err:=db.DynamicSelect(&users,"select * from public.user",whereMap3,nil,"desc",2,0,arg1,arg2,arg3)
            	if err!=nil{
            		fmt.Println(err)
            		return 
            	}
    	    
        	//4.2DynamicSelectOne
			    var user = User{}
        	    arg1 :="ft"
				arg2 :=1035
				arg3 :=1
            	orderBy3 :=make([]string,0)
            	orderBy3 = append(orderBy3,"id")
            	whereMap3 := make([][]string,0)
				
				if arg1!=""{
				    whereMap3 = append(whereMap3,[]string{
            		"","name","=",
            	    })
				}

				if arg2!=0{
				    whereMap3 = append(whereMap3,[]string{
            		"","age","=",
            	    })
				}
				if arg3!=0{
				    whereMap3 = append(whereMap3,[]string{
            		"","class","=",
            	    })
				}				
            	count,err:=db.DynamicSelectOne(&user,"select * from public.user",whereMap3,nil,"desc",2,0,arg1,arg2,arg3)
            	if err!=nil{
            		fmt.Println(err)
            		return 
            	}
            
            //4.3 DynamicSelectCount   // 弃用,推荐使用DynamicSelectOne(&count,"select count(*) from xx",x,x,x,x,x,x)
			    arg1 :="ft"
				arg2 :=1035
				arg3 :=1
            	orderBy3 :=make([]string,0)
            	orderBy3 = append(orderBy3,"id")
            	whereMap3 := make([][]string,0)
				
				if arg1!=""{
				    whereMap3 = append(whereMap3,[]string{
            		"","name","=",
            	    })
				}

				if arg2!=0{
				    whereMap3 = append(whereMap3,[]string{
            		"","age","=",
            	    })
				}
				if arg3!=0{
				    whereMap3 = append(whereMap3,[]string{
            		"","class","=",
            	    })
				}				
            	count,err:=db.DynamicSelectCount("select * from public.user",whereMap3,nil,"desc",2,0,arg1,arg2,arg3)
            	if err!=nil{
            		fmt.Println(err)
            		return 
            	}
            	fmt.Println(count)
		//5. 动态连表查询 DynamicJoinSelect/DynamicJoinSelectOne,区别同上
		    	orders := make([]PayOrder, 0)
				orderBy := make([]string, 0)
				orderBy = append(orderBy, "ttime")


				joinMap :=make([]string,0)
				joinMap = append(joinMap,"o.productid=a.productId")
				whereMap := make([][]string, 0)
				whereMap = append(whereMap, []string{
					"and", "a.adsId", "=",
				})
				//select distinct o.* from productid_ads a,payorder o where o.productid=a.productId and a.adsId='666@qq.com'

				err := DynamicJoinSelect(&orders, "select distinct o.* from productid_ads a,payorder o", whereMap,joinMap, orderBy, "desc", 20, 0, "666@qq.com")
				if err != nil {
					t.Fatal(err)
				}
				t.Log(orders)
		//6. 多数据源操作**跨源操作仅仅能使用xorm原生db**
				db.NewDataSource("source2","xxxxxxx")
				db.NewDataSource("source3","xxxxxxxx")
				使用:
				var engine *xorm.Engine = db.Dbs["source2"]
				engine.SQL().Get()
				engine.SQL().Find()
				engine.Exec()
				用法同xorm的engine对象。
		//7. 全局session
		     session :=db.LocalSession
			 多源:
			 session :=db.LocalSessions["source2"]
			 session用法同engine
				session.SQL().Get()
				session.SQL().Find()
				session.Exec()
		
		//8.动态update
		  type Tmp struct{
		  	Name string `xorm:"name"`
		  	Id int `xorm:"id"`
		  }
		  t:= Tmp{Name:"高中m班",Id:7}
		  whereMap := make([][]string,0)
		  var teacherName = "Mr X"
		  if teacherName!=""{
		  	whereMap = append(whereMap,[]string{
		  		"","teachername","=",
		  	})
		  }
		  db.DynamicUpdate("default","update class",t,whereMap,t.Name,t.Id,teacherName)
		  
	}	  
```