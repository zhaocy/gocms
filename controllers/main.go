package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Index() {
	this.TplName = "public/main.html"
}
