package controllers

import (
	"github.com/astaxie/beego"
	"freecake/models"
)

type HouseIndexController struct {
	beego.Controller
}

func (this *HouseIndexController) retData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()

}
func (this *HouseIndexController) GetHouseIndex(){
	resp := make(map[string]interface{})

	resp["errno"] = models.RECODE_DATAERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
	this.retData(resp)
}