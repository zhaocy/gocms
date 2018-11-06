package controllers

import (
	"github.com/astaxie/beego"
	"gocms/libs"
	"gocms/models"
	"strconv"
	"strings"
	"time"
)

type LoginController struct {
	BaseController
}


func (self *LoginController) Login(){
	beego.ReadFromRequest(&self.Controller)
	if self.isPost() {

		username := strings.TrimSpace(self.GetString("username"))
		password := strings.TrimSpace(self.GetString("password"))

		if username != "" && password != "" {
			user, err := models.AdminGetByName(username)
			flash := beego.NewFlash()
			errorMsg := ""
			if err != nil || user.Password != libs.Md5([]byte(password+user.Salt)) {
				errorMsg = "帐号或密码错误"
			} else if user.Status == -1 {
				errorMsg = "该帐号已禁用"
			} else {
				user.LastIp = self.getClientIp()
				user.LastLogin = time.Now().Unix()
				user.Update()
				authkey := libs.Md5([]byte(self.getClientIp() + "|" + user.Password + user.Salt))
				self.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)

				self.redirect(beego.URLFor("HomeController.Index"))
			}
			flash.Error(errorMsg)
			flash.Store(&self.Controller)
			self.redirect(beego.URLFor("LoginController.LoginIn"))
		}
	}
	self.TplName = "login/login.html"
}


type User struct {
	Roles []string `json:"roles"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
}

//{"code":20000,"data":{"roles":["admin"],"name":"admin","avatar":"https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"}}
func (self *LoginController) UserInfo(){
	//self.Ctx.Request.Header.Set("Access-Control-Allow-Origin", "*")
	out := make(map[string]interface{})
	out["code"] = 20000
	user := User{
		Roles:[]string{"admin"},
		Name:"admin",
		Avatar: "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
	}
	out["data"] = user
	self.Data["json"] = out
	self.ServeJSON()
}
