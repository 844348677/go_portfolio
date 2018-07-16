package controllers

import (
	"github.com/astaxie/beego"
	"freecake/models"
	"github.com/astaxie/beego/orm"

)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) retData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()

}

func (c *AreaController) GetArea() {

	beego.Info("connect success")
	beego.Info("connect success")
	resp := make(map[string]interface{})

	//无论怎样都要最后执行
	defer c.retData(resp)
	//从ｓｅｓｓｉｏｎ拿数据

	//从数据库拿到数据
	var areas []models.Area

	o := orm.NewOrm()
	num,err := o.QueryTable("area").All(&areas)
	if err != nil{
		beego.Info("数据错误")
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

		return
	}
	if num == 0{
		beego.Info("没有数据")
		resp["errno"] = 4002
		resp["errmsg"] = "没有查到数据"

		return
	}



	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = areas
	beego.Info("query data success , data = " ,resp)
	//打包成　ｊｓｏｎ　返回给前端

}
