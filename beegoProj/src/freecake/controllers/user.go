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
	this.SetSession("user_id",user.Id)
	this.SetSession("mobile",user.Mobile)

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
	//路径要全
	urlMap["avatar_url"] = "http://192.168.1.104:8080/" + uploadResponse.RemoteFileId
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = urlMap

}

//GetUserData
func (this *UserController) GetUserData() {
	resp := make(map[string]interface{})
	defer this.retData(resp)

	//1 从session 里面获取　userid
	user_id := this.GetSession("user_id")


	//2 从数据库中拿到　user_id 对应的 user值
	// GetSession 万能接口类型　需要断言
	user := models.User{Id:user_id.(int)}
	o := orm.NewOrm()
	err := o.Read(&user)
	if err != nil{
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = &user


}

func (this *UserController) UpdateName(){
	resp := make(map[string]interface{})
	defer this.retData(resp)

	//1 获得　session 中的user_id
	user_id := this.GetSession("user_id")

	//2 获取前段传过来的数据
	// beego　文档　请求数据处理
	UserName := make(map[string]string)
	json.Unmarshal(this.Ctx.Input.RequestBody,&UserName)
	beego.Info("get userName = ",UserName["name"])

	//3 更新　user_id 对应的name
	o := orm.NewOrm()
	user := models.User{Id:user_id.(int)}
	if err := o.Read(&user); err != nil{
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	user.Name = UserName["name"]
	beego.Info("data user = ",user)

	if _,err := o.Update(&user); err != nil{
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	//4 把session中的name字段更新
	this.SetSession("name",UserName["name"])

	//5 把数据打包返回给前段
	resp["data"] = UserName
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

}

func (this *UserController) PostRealName(){
	resp := make(map[string]interface{})
	defer this.retData(resp)

	// session 获得　user_id
	user_id := this.GetSession("user_id")

	// 获取前段穿过来的数据
	realName := make(map[string]string)
	json.Unmarshal(this.Ctx.Input.RequestBody,&realName)
	beego.Info("get realName = ",realName["real_name"])
	beego.Info("get id_card = ",realName["id_card"])

	// 更新数据库中　userId对应的表的信息
	o := orm.NewOrm()
	user := models.User{Id:user_id.(int)}

	if err := o.Read(&user); err != nil{
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	// 设置　需要插入的信息
	user.Real_name = realName["real_name"]
	user.Id_card = realName["id_card"]

	if _,err := o.Update(&user); err != nil{
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	this.SetSession("user_id",user.Id)
	//打包　json数据返回给前段
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)


}