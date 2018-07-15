package routers

import (
	"helloworld/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
    //正则表达式 ?一个或多个字符　*星能匹配０个
    //不能有选择困难
	beego.Router("/user/?:id", &controllers.UserController{},"get:GetInfo;post:PostData")

}
