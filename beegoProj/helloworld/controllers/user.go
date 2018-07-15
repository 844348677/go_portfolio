package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString( "helloworld\n")
	c.Ctx.WriteString("getinfo responce success")
}

func (c *UserController) GetInfo(){

	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("getInfo data ,id = "+id)

}

//curl -d "sssss" "127.0.0.1:8080/user/1111"
func (c *UserController) PostData(){
	c.Ctx.WriteString("this is post func\n")
}

