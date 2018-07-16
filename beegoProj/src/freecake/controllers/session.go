package controllers

import (
	"github.com/astaxie/beego"
	"freecake/models"
)

type SessionController struct {
	beego.Controller
}

func (this *SessionController) retData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()

}
func (this *SessionController) GetSessionData(){
	resp := make(map[string]interface{})
	defer this.retData(resp)
	user := models.User{}
	//user.Name = "liuh"

	resp["errno"] = models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

	name := this.GetSession("name")
	if name != nil{
		user.Name = name.(string)
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"] = user
	}

}

//session sessionID hash cookie
//增加一个session　模块
//在项目里面　启用session 两种设置　config main()函数
//要在注册结束后去增加session内容 在首页判断session是否有值

func (this *SessionController) DeleteSessionData(){
	resp := make(map[string]interface{})
	defer this.retData(resp)
	this.DelSession("name")
	
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}
