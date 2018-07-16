package main

import (
	_ "freecake/routers"
	"github.com/astaxie/beego"
	"strings"
	"net/http"
	"github.com/astaxie/beego/context"
	_ "freecake/models"
)
// setting -> go -> gopath -> project gopath (src的上层目录)

func main() {
	ignoreStaticPath()
	beego.Run()
}

func ignoreStaticPath(){
	beego.InsertFilter("/",beego.BeforeRouter,TransparentStatic)
	beego.InsertFilter("/*",beego.BeforeRouter,TransparentStatic)

}
func TransparentStatic(ctx *context.Context){
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ",orpath)
	if strings.Index(orpath,"api") >=0{
		return
	}
	http.ServeFile(ctx.ResponseWriter,ctx.Request,"static/html/"+ctx.Request.URL.Path)
}
