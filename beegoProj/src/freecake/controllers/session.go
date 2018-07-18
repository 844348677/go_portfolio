package controllers

import (
	"github.com/astaxie/beego"
	"freecake/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
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

func (this *SessionController) Login(){
	//得到用户信息
	resp := make(map[string]interface{})
	defer this.retData(resp)

	//获取前段传过来的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody,&resp)
	//beego.Info("========name = ",resp["mobile"]," =========password = ",resp["password"])
	//判断是否合法
	if resp["mobile"] == nil || resp["password"] == nil{
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}


	//与数据库内容匹配　账号密码正确
	o := orm.NewOrm()
	user := models.User{Name:resp["mobile"].(string)}
	qs := o.QueryTable("user")
	_,err := qs.Filter("mobile",resp["mobile"].(string)).All(&user)
	if err != nil{
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}

	if user.Password_hash != resp["password"]{
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}

	//添加session　
	this.SetSession("name",user.Name)
	this.SetSession("mobile",resp["mobile"])
	this.SetSession("user_id",user.Id)

	//返回json数据给前段
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

}
