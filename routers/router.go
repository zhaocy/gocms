package routers

import (
	"gocms/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.WwwController{},"*:Index")
	beego.Router("/show/:id",&controllers.WwwController{},"*:Show")
    beego.Router("/list/:class_id",&controllers.WwwController{},"*:List")

    beego.Router("/login", &controllers.LoginController{},"*:Login")

	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")
	beego.Router("/article/list", &controllers.ArticleController{}, "*:List")

    beego.AutoRouter(&controllers.ArticleController{})
}
