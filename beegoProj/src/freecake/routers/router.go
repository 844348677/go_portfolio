package routers

import (
	"freecake/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1.0/areas", &controllers.AreaController{},"get:GetArea")
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{},"get:GetHouseIndex")
	beego.Router("/api/v1.0/session", &controllers.SessionController{},"get:GetSessionData;delete:DeleteSessionData")
	beego.Router("/api/v1.0/users", &controllers.UserController{},"post:Reg")
    //重复　覆盖问题
	//beego.Router("/api/v1.0/session", &controllers.UserController{},"delete:Reg")
	beego.Router("/api/v1.0/sessions", &controllers.SessionController{},"post:Login")
	// /api/v1.0/user/avatar
	beego.Router("/api/v1.0/user/avatar", &controllers.UserController{},"post:PostAvatar")
    // /api/v1.0/user
	beego.Router("/api/v1.0/user", &controllers.UserController{},"get:GetUserData")
    // /v1.0/user/name
	beego.Router("/api/v1.0/user/name", &controllers.UserController{},"put:UpdateName")
	// /v1.0/user/auth
	beego.Router("/api/v1.0/user/auth", &controllers.UserController{}, "get:GetUserData;post:PostRealName")

}
