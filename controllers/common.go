package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	noLayout       bool
	userId         int
	userName       string
	loginName      string
	pageSize       int
}

//在调用路由之前执行
func (self *BaseController) Prepare() {
	controllerName, actionName := self.GetControllerAndAction()
	//fmt.Printf("%s - %s \n", controllerName, actionName)
	self.controllerName = strings.ToLower(controllerName[:len(controllerName)-10])
	self.actionName = strings.ToLower(actionName)
	self.pageSize = 16
}

/*
加载模板
 */
func (this *BaseController) display(tpl ...string) {

	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = this.controllerName + "/" + this.actionName + ".html"
	}

	if !this.noLayout {
		if this.Layout == "" {
			this.Layout = "public/layout.html"
		}
	}
	fmt.Printf("TplName = %s\n", tplname)
	this.TplName = tplname
}

// 是否POST提交
func (self *BaseController) isPost() bool {
	return self.Ctx.Request.Method == "POST"
}

// 重定向
func (self *BaseController) redirect(url string) {
	self.Redirect(url, 302)
	self.StopRun()
}

//获取用户IP地址
func (self *BaseController) getClientIp() string {
	s := strings.Split(self.Ctx.Request.RemoteAddr, ":")
	return s[0]
}