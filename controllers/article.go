package controllers

import (
	"gocms/models"
)

type ArticleController struct {
	BaseController
}

func (self *ArticleController)List(){

}


type ArticleData struct {
	Items []*models.Article `json:"items"`
}

/*
{
  "data" : {
    "items" : [
      {
        "status" : "published",
        "id" : "620000201003044567",
        "title" : "Fyth gtkd fktxy uicbc jfhq xtqj fblhcc smo xkpp srvc mzvycb zdytou uxzhsirsfo ofgttr lxheee psqtw jyojtonna ukisrzkbpn.",
        "author" : "name",
        "pageviews" : 4204,
        "display_time" : "1972-04-11 09:39:11"
      },
	...
	]
	}
}
 */
func (self *ArticleController)Table(){
	out := make(map[string]interface{})
	list := models.ArticleGetList()
	item := ArticleData{
		Items:list,
	}
	out["data"] = item
	self.Data["json"] = out
	self.ServeJSON()
}
