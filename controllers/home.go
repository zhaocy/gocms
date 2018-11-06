package controllers


type HomeController struct {
	BaseController
}

func (this *HomeController) Start(){
	this.Data["pageTitle"] = "控制面板"
	this.display()
}