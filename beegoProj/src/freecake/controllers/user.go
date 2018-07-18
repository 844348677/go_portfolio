package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"freecake/models"
	"path"
	"github.com/weilaihui/fdfs_client"
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

//PostAvatar
func (this *UserController) PostAvatar(){
	//拿到二进制数据
	resp := make(map[string]interface{})
	defer this.retData(resp)

	//１ 获取上传的一个文件
	fileData,hd,err := this.GetFile("avatar")
	if err != nil{
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}
	//2 得到文件后缀
	// beego path包　ext()　截取后缀函数
	suffix := path.Ext(hd.Filename) // a.jpg.mp3

	//3 存储文件到　fastdfs上
	fdfsClient,err := fdfs_client.NewFdfsClient("/home/liuh/go/src/freecake/conf/client.conf")
	if err != nil{
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}
	fileBuffer := make([]byte,hd.Size)
	_,err = fileData.Read(fileBuffer)
	if err != nil{
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}
	// suffix 获得的是.jpg 拼接多了一个点　　xxxxxx..jpg
	uploadResponse,err := fdfsClient.UploadByBuffer(fileBuffer,suffix[1:])
	if err != nil{
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}
	//uploadResponse.GroupName
	//uploadResponse.RemoteFileId

	//4 从session 里面拿到user_id
	user_id := this.GetSession("user_id")
	var user models.User
	//5 更新用户　数据库中的内容　freecake库　user表中　有一个　avatar_url的字段
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	qs.Filter("Id",user_id).One(&user)
	user.Avatar_url = uploadResponse.RemoteFileId

	_,errUpdate := o.Update(&user)
	if errUpdate != nil{
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	//Avaurl := "192.168.1.104:8080/" + uploadResponse.RemoteFileId

	urlMap := make(map[string]string)
	urlMap["avatar_url"] = "192.168.1.104:8080/" + uploadResponse.RemoteFileId
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = urlMap

}