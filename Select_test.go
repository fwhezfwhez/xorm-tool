package xormTool

import (
	"testing"
	"time"
)

func TestSelectCount(t *testing.T) {
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	count, err := SelectCount("select * from class limit ? offset ?",10,1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(count)
}

func TestSelect(t *testing.T) {
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	type Class struct {
		Id   int
		Name string
	}
	classes := make([]Class, 0)
	err := Select(&classes, "select * from class")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(classes)

}
func TestSelectOne(t *testing.T) {
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	type Class struct {
		Id   int
		Name string
	}
	class := Class{}
	err := SelectOne(&class, "select * from class where id=?", 5)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(class)
}

func TestDynamicSelect(t *testing.T) {
	type User struct {
		Id          int
		Name        string
		ClassId     int `db:"class_id"`
		Description string
	}
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	users := make([]User, 0)
	orderBy := make([]string, 0)
	orderBy = append(orderBy, "id")
	whereMap := make([][]string, 0)
	whereMap = append(whereMap, []string{
		"", "name", "=",
	})
	whereMap = append(whereMap, []string{
		"and", "id", ">",
	})
	whereMap = append(whereMap, []string{
		"and", "id", "<>",
	})
	whereMap = append(whereMap, []string{
		"and", "created", ">",
	})
	err := DynamicSelect(&users, "select * from public.user", whereMap, orderBy, "desc", 2, 0, "ft4", "1035", "1","2012/09/09")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(users)
}

func TestDynamicSelectOne(t *testing.T) {
	type User struct {
		Id          int
		Name        string
		ClassId     int `db:"class_id"`
		Description string
	}
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	user := User{}
	orderBy := make([]string, 0)
	orderBy = append(orderBy, "id")
	whereMap := make([][]string, 0)
	whereMap = append(whereMap, []string{
		"", "id", "=",
	})
	err := DynamicSelectOne(&user, "select * from public.user", whereMap, nil, "desc", 2, 0, 1036)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestDynamicSelectCount(t *testing.T) {
	type User struct {
		Id          int
		Name        string
		ClassId     int `db:"class_id"`
		Description string
	}
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	orderBy := make([]string, 0)
	orderBy = append(orderBy, "id")
	whereMap := make([][]string, 0)
	whereMap = append(whereMap, []string{
		"", "name", "=",
	})
	whereMap = append(whereMap, []string{
		"and", "id", ">",
	})
	whereMap = append(whereMap, []string{
		"and", "id", "<>",
	})
	count, err := DynamicSelectCount("select * from public.user", whereMap, nil, "desc", 2, 0, "ft4", 1035, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(count)
}

func TestRemove(t *testing.T) {
	ti, er := time.Parse("2006-01-02 15:04:05", "2018-04-13 15:30:33")
	if er != nil {
		t.Fatal(er)
	}
	args := []interface{}{
		1, ti,1,
	}
	t.Log("原始数据:",args)
	args = Remove(args, 1)
	t.Log("删除后:",args)
}
func BenchmarkRemove(b *testing.B){
	ti, er := time.Parse("2006-01-02 15:04:05", "2018-04-13 15:30:33")
	if er != nil {
		b.Fatal(er)
	}
	args := []interface{}{
		1, ti,1,
	}
	b.Log("原始数据:",args)
	for i:=0;i<b.N;i++{
		args = Remove(args, 1)
	}
	b.Log("删除后:",args)
}
func TestRemoveZero(t *testing.T) {
	ti, er := time.Parse("2006-01-02 15:04:05", "2018-04-13 15:30:33")
	if er != nil {
		t.Fatal(er)
	}
	args := []interface{}{
		1, 0.5, "test", ti,"","%%",
	}
	args=RemoveZero(args)
	t.Log(args)
}

func BenchmarkRemoveZero(b *testing.B){
	ti, er := time.Parse("2006-01-02 15:04:05", "2018-04-13 15:30:33")
	args := []interface{}{
		1, 0.5, "test", ti,"","%%",
	}
	if er != nil {
		b.Fatal(er)
	}
	for i:=0;i<b.N;i++{
		args=RemoveZero(args)
	}
	b.Log(args)
}

func TestDynamicJoinSelect(t *testing.T){
	type PayOrder struct {
		Id          int       `xorm:"id" json:"id"`
		OrderId     string    `xorm:"orderid" json:"order_id"`
		Cid         string    `xorm:"cid" json:"cid"`
		Pversion    string    `xorm:"pversion" json:"pVersion"`
		Ttime       time.Time `xorm:"ttime" json:"tTime"`
		Ip          string    `xorm:"ip" json:"ip"`
		OrderStatus string    `xorm:"orderstatus" json:"order_status"`
		Phone       string    `xorm:"phone" json:"phone"`
		Uname       string    `xorm:"uname" json:"uName"`
		PayStatus   string    `xorm:"paystatus" json:"pay_status"`
		Money       string    `xorm:"money" json:"money"`
		BillType    string    `xorm:"billtype" json:"bill_type"`
		ProductId   string    `xorm:"productid" json:"product_id"`
		TtimeUnix   string    `xorm:"ttimeunix" json:"tTime_unix"`

		TravelDate string `xorm:"traveldate" json:"travel_date"`
		Num        string `xorm:"num" json:"num"`
		UserNames  string `xorm:"usernames" json:"user_names"`
		AppType    string `xorm:"apptype" json:"app_type"`
	}


	DataSource("postgres://medium:mediuml4eLxglxL8@111.231.137.127:5432/medium?sslmode=disable")
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
}

func TestPrepareIn(t *testing.T) {
	var args = []interface{}{
		1,2,3,4,5,
	}
	t.Log(PrepareIn(args))
}

func TestSelectIn(t *testing.T) {
	DataSource("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	Db.ShowSQL(true)
	type Class struct {
		Id   int
		Name string
	}
	var args = []interface{}{1,2,3,4,5,6,7,8,9,10}
	classes := make([]Class, 0)
	whereMap :=make([][]string,0)
	whereMap = append(whereMap,[]string{
		"","id","in",PrepareIn(args),
	})
	err := DynamicSelect(&classes, "select * from class",whereMap,nil,"",-1,-1,args...)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(classes)

}