package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"gocms/models"
	"math/rand"
	"time"
)

type WwwController struct {
	BaseController
}

func (self *WwwController) Index() {
	//创建切片，过滤条件
	filter := make([]interface{}, 0)
	filter = append(filter, "status", 1)
	filter = append(filter, "class_id", 2)

	result, _ := models.NewsGetList(1, 10, filter...)

	//获取诗词古韵中的前六条记录
	list := make([]map[string]interface{}, len(result))

	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		if v.Picurl == "" {
			var r = rand.Intn(16)
			v.Picurl = "/upload/rand" + fmt.Sprintf("%d", r) + ".jpeg"
		}
		row["picurl"] = v.Picurl
		row["media"] = v.Media
		row["title"] = v.Title
		if v.Desc != "" {
			runes := []rune(v.Desc)
			length := len(runes)
			if length > 30 {
				length = 30
			}
			row["desc"] = string(runes[:length])
		}
		list[k] = row
	}

	out := make(map[string]interface{})

	out["list"] = list
	out["class_id"] = 0
	self.Data["data"] = out
	self.Layout = "public/www_layout.html"
	self.display()

}

/*
 * /show/:id
 */
func (self *WwwController) Show() {
	id, _ := self.GetInt(":id")
	news, _ := models.NewsGetById(id)
	row := make(map[string]interface{})
	row["class_id"] = 0
	if news != nil {
		row["title"] = news.Title
		if news.Picurl == "" {
			var r = rand.Intn(16)
			news.Picurl = "/upload/image/rand" + fmt.Sprintf("%d", r) + ".jpeg"
		}
		row["picurl"] = news.Picurl
		row["media"] = news.Media
		row["content"] = news.Content
	}
	self.Data["data"] = row
	self.Layout = "public/www_layout.html"
	self.display()
}

//
func (self *WwwController) List() {

	page, err:= self.GetInt("page")
	if err != nil {
		page = 1
	}
	pagesize := self.pageSize
	catId, err := self.GetInt(":class_id")
	filters := make([]interface{}, 0)
	if err == nil {
		filters = append(filters, "class_id", catId)
	}
	filters = append(filters, "status", 1)
	result, count := models.NewsGetList(page, pagesize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		if v.Picurl == "" {
			var r = rand.Intn(16)
			v.Picurl = "/upload/image/rand" + fmt.Sprintf("%d", r) + ".jpeg"
		}
		row["picurl"] = v.Picurl
		row["media"] = v.Media
		row["title"] = v.Title
		if v.Desc !=""{
			nameRun:= []rune(v.Desc)
			length:=len(nameRun)
			if length>30 {
				length = 30
			}
			row["desc"] = string(nameRun[:length])
		}
		row["author"] = v.Author
		row["posttime"] = beego.Date(time.Unix(v.Posttime, 0), "Y-m-d")
		list[k] = row

	}
	classArr := make(map[int]string)
	classArr[5] = "开心儿歌"
	classArr[3] = "儿童古诗"
	classArr[2] = "诗词古韵"
	classArr[1] = "经典国学"

	out := make(map[string]interface{})
	out["count"] = count
	out["class_id"] = catId
	out["page"] = page
	out["class_name"] = classArr[catId]
	out["title"] = classArr[catId]
	out["list"] = list
	self.Data["data"] = out

	self.Layout = "public/www_layout.html"
	self.display()
}
