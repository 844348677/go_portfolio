package models

import (
	 _ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

//注意顺序　注意大小写　名字和类型名一样
type User struct {
	Id   int //`orm:pk,auto`//主键　自增　
	Name string `orm:"size(100)"` //varchar(100)

	User_order []*User_order `orm:"reverse(many)"`
}
type User_order struct{
	Id int
	Order_data string `orm:size(100)`

	User *User `orm:"rel(fk)"`
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/testbeego?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User),new(User_order))

	// create table
	orm.RunSyncdb("default", false, true)
}


