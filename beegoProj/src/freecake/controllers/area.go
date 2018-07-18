package controllers

import (
	"github.com/astaxie/beego"
	"freecake/models"
	"github.com/astaxie/beego/orm"
	//插入　beego 的　redis
	_ "github.com/astaxie/beego/cache/redis"
	//_ "github.com/gomodule/redigo/redis"
	"github.com/astaxie/beego/cache"
	"time"
	"fmt"
	"encoding/json"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) retData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()

}

func (c *AreaController) GetArea() {

	beego.Info("connect success")
	beego.Info("connect success")
	resp := make(map[string]interface{})

	//无论怎样都要最后执行
	defer c.retData(resp)
	//从ｓｅｓｓｉｏｎ拿数据


	//https://github.com/golang
	//conn refuse

	// go get github.com/astaxie/beego/cache
	// go get -u github.com/astaxie/beego/cache/memcache
	// import _ "github.com/astaxie/beego/cache/memcache"
	// go get github.com/garyburd/redigo/redis
	// go get github.com/gomodule/redigo/redis
	//　在／go/src/github.com/astaxie/beego/cache/redis/redis.go　中修改import包名
	//从redis缓存拿数据　网站的优化
	// redis; go get github.com/astaxie/beego/cache; go get github.com/gomodule/redigo/redis
	// redis.conf daemonize yes 后台执行
	//　查看端口是否被占用　netstat -ano | grep 6379

	//　程序启动　需要开启　redis　服务器
	// sudo /home/liuh/redis/redis-4.0.10/src/redis-server ./redis.conf
	// /home/liuh/redis/redis-4.0.10/src/redis-cli -p 6379

	// select 1
	// keys *
	// set "aa" "bb"
	// key *
	// get "aa"
	cache_conn, err := cache.NewCache("redis", `{"key":"freecake","conn":":6379","dbNum":"0"}`)
	//注意有freecake:area
	if areaData := cache_conn.Get("area"); areaData != nil{
		beego.Info("get data from cache ============")
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		//fmt.Printf("%s \n",areaData)
		str := fmt.Sprintf("%s",areaData)
		//fmt.Println(str)
		//fmt.Println(strings.Contains(str,`\`))
		var areas []models.Area
		err1 := json.Unmarshal([]byte(str),&areas)
		if err1 != nil{
			beego.Info("err11111111 : ",err1)
		}
		//beego.Info("err11111111 : ",areas)
		resp["data"] = areas

		return
	}

	// freecake 类似namespace
	//errCache := cache_conn.Put("aaa","bbb",time.Second*60)
	//if errCache != nil{
	//	beego.Error("cache err")
	//}
	//beego.Info("cache_conn.aaa = ",cache_conn.Get("aaa"))
	//二进制打印
	//fmt.Printf("cache_conn , conn[a] = %s\n",cache_conn.Get("aaa"))

	//从数据库拿到数据
	var areas []models.Area

	o := orm.NewOrm()
	num,err := o.QueryTable("area").All(&areas)
	if err != nil{
		beego.Info("数据错误")
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

		return
	}
	if num == 0{
		beego.Info("没有数据")
		resp["errno"] = 4002
		resp["errmsg"] = "没有查到数据"

		return
	}

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = areas

	//把数据转换成json格式　存入缓存中
	json_str,err := json.Marshal(areas)

	if err!= nil{
		beego.Info("encoding err = ",err)
		return
	}
	beego.Info("json_str ====== ",json_str)
	cache_conn.Put("area",json_str,time.Second*3600)


	//打包成　ｊｓｏｎ　返回给前端
	beego.Info("query data success , data = " ,resp)


}
