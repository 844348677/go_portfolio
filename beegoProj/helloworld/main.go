package main

import (
	_ "helloworld/routers"
	"github.com/astaxie/beego"
	_ "helloworld/models"
	"github.com/astaxie/beego/orm"
	"helloworld/models"
)

func inserUser(){
	o := orm.NewOrm()
	user := models.User{}
	user.Name = "liuh"
	id,err := o.Insert(&user)
	if err != nil{
		beego.Info("insert error = ",err)
		return
	}
	beego.Info("insert success, id = ",id)
}

func queryUser(){
	o := orm.NewOrm()
	user := models.User{Id:1}
	err := o.Read(&user)
	if err != nil{
		beego.Info("query error ",err)
		return
	}
	beego.Info("query sucess , user = ",user)
}

func userUpdate(){
	o := orm.NewOrm()
	user := models.User{Id:1,Name:"liuhang"}
	_,err := o.Update(&user)
	if err != nil{
		beego.Info("query error ",err)
		return
	}
	beego.Info("update sucess , user = ",user)
}

func deleteUser(){
	o := orm.NewOrm()
	user := models.User{Id:2}
	_,err := o.Delete(&user)
	if err != nil{
		beego.Error("delete error = ",err)
		return
	}
	beego.Info("delete success, id = ")
}

func insertorder(){
	o := orm.NewOrm()
	order := models.User_order{}
    order.Order_data = "this is order"
    user := models.User{Id:1}
    order.User = &user

	id,err := o.Insert(&order)
	if err != nil{
		beego.Info("insert error = ",err)
		return
	}
	beego.Info("insert success, id = ",id)
}

func queryOrder(){
	var orders []*models.User_order
	o := orm.NewOrm()
	qs := o.QueryTable("User_order")
	_,err := qs.Filter("user_id",1).All(&orders)

	if err != nil{
		beego.Info("query error = ",err)
		return
	}
	for _,order := range orders{
		beego.Info("insert success, v = ",order)
	}
}


func main() {
	//inserUser()
	//queryUser()
	//userUpdate()
	//deleteUser()
	//insertorder()
	queryOrder()
	//beego.SetStaticPath("download","down")
	beego.Run()

}

