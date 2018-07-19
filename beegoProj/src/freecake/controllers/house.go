package controllers

import (
	"github.com/astaxie/beego"
	"freecake/models"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"strconv"
)

type HouseController struct {
	beego.Controller
}

func (this *HouseController) retData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()

}

func (this *HouseController) GetHouseData(){
	resp := make(map[string]interface{})
	defer this.retData(resp)

	//1 从session 里面获取　userid
	user_id := this.GetSession("user_id")

	// select * from house where user.id == user_id
	houses := []models.House{}

	o := orm.NewOrm()
	qs := o.QueryTable("house")

	// num,err := qs.Filter("user__id",user_id.(int)).All(&houses)
	num,err := qs.Filter("user_id",user_id.(int)).All(&houses)
	if err != nil{
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	if num == 0{
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}
	respData := make(map[string]interface{})
	respData["houses"] = houses

	resp["data"] = &respData
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

}

func (this *HouseController) PostHouseData() {
	resp := make(map[string]interface{})
	defer this.retData(resp)
	// 1 从前段拿到数据
	reqData := make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody,&reqData)

	// 2 判断前段数据的合法性


	// 3 插入数据到数据库
	house := models.House{}

	house.Title = reqData["title"].(string)
	price,_ := strconv.Atoi(reqData["price"].(string))
	house.Price = price
	house.Address = reqData["address"].(string)
	room_count,_ := strconv.Atoi(reqData["room_count"].(string))
	house.Room_count = room_count
	acreage,_ := strconv.Atoi(reqData["acreage"].(string))
	house.Acreage = acreage
	house.Unit = reqData["unit"].(string)
	capacity,_ := strconv.Atoi(reqData["capacity"].(string))
	house.Capacity = capacity
	house.Beds = reqData["beds"].(string)
	deposit,_ := strconv.Atoi(reqData["deposit"].(string))
	house.Deposit = deposit
	min_days,_ := strconv.Atoi(reqData["min_days"].(string))
	house.Min_days = min_days
	max_days,_ := strconv.Atoi(reqData["max_days"].(string))
	house.Max_days = max_days

	facilities := []models.Facility{}
	for _,data := range reqData["facility"].([]interface{}){
		f_id,_ := strconv.Atoi(data.(string))
		fac := models.Facility{Id:f_id}
		facilities = append(facilities, fac)
	}

	area_id,_ := strconv.Atoi(reqData["area_id"].(string))
	area := models.Area{Id:area_id}
	house.Area = &area
	user_id := this.GetSession("user_id").(int)
	user := models.User{Id:user_id}
	house.User = &user

	o := orm.NewOrm()
	// 加指针 house_id,err := o.Insert(&house)
	_,err := o.Insert(&house)
	if err != nil{
		beego.Info("11111111111111")
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	// ???
	//house.Id = int(house_id)

	// models.House 结构中的field  Facilities
	m2m := o.QueryM2M(&house,"Facilities")
	num,errM2M := m2m.Add(facilities)
	if errM2M != nil || num == 0{
		beego.Info("222222222222 ",num)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	respData := make(map[string]interface{})
	respData["house_id"] = house.Id
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}

func (this *HouseController) GetDetailHouseData() {
	resp := make(map[string]interface{})
	defer this.retData(resp)

	respData := make(map[string]interface{})

	// 1 获取当前用户的　user_id

	user_id := this.GetSession("user_id")

	// 2 从url中获取房屋　id
	//beego文档 路由设置　正则路由
	house_id := this.Ctx.Input.Param(":id")

	// 3 从缓存中获取当前房屋数据　redis

	// 4 关联查询　
	o := orm.NewOrm()
	h_id,err := strconv.Atoi(house_id)
	if err != nil{
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}
	house := models.House{Id:h_id}
	// 需要读数据
	if err := o.Read(&house) ; err != nil{
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	/*
	respData["acreage"] = house.Acreage
	respData["address"] = house.Address
	respData["beds"] = house.Beds
	respData["capacity"] = house.Capacity
	respData["deposit"] = house.Deposit
	respData["facilities"] = house.Facilities
	respData["img_urls"] = house.Images
	respData["max_days"] = house.Max_days
	respData["min_days"] = house.Min_days
	respData["price"] = house.Price
	respData["room_count"] = house.Room_count
	*/

	o.LoadRelated(&house,"Area")
	o.LoadRelated(&house,"User")
	o.LoadRelated(&house,"Images")
	o.LoadRelated(&house,"Facilities")
	beego.Info("respData = ",respData)

	user := models.User{Id:user_id.(int)}
	house.User = &user

	var facilities []int
	for _,data := range house.Facilities{
		facilities = append(facilities, data.Id)
	}
	respData["facilities"] = facilities

	respData["house"] = house

	//o.QueryTable("")

	// 5 存入缓存　

	//  6 打包返回json

	resp["data"] = respData
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

}