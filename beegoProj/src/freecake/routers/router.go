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

}
