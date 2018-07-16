package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"freecake/models"
)
//设置路由
//添加user.go　写user.go代码
//先设置配置文件


type UserController struct {
	beego.Controller
}

func (this *UserController) retData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()

}
//在配置文件设置　copyrequestbody = true
func (this *UserController) Reg(){
	resp := make(map[string]interface{})
	defer this.retData(resp)

	//获取前段传过来的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody,&resp)

	//beego.Info(`resp["mobile"] = `,resp["mobile"])
	//beego.Info(`resp["password"] = `,resp["password"])
	//beego.Info(`resp["sms_code"] = `,resp["sms_code"])
	o := orm.NewOrm()
	user := models.User{}
	user.Password_hash = resp["password"].(string)
	user.Name = resp["mobile"].(string)
	user.Mobile = resp["mobile"].(string)

	id,err := o.Insert(&user)
	if err != nil{
		resp["errno"] = 4002
		resp["errmsg"] = "注册失败"
		return
	}
	beego.Info("reg success , id = ",id)
	resp["errno"] = 0
	resp["errmsg"] = "注册成功"

	this.SetSession("name",user.Name)

}