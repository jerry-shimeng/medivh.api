package controllers
import (
	"github.com/astaxie/beego"
	"rwapi/code"
)

type StatusControllers struct {
	beego.Controller
}

//连接状态
//@router /all [get]
func (this *StatusControllers)GetAll(){
	r := code.StatusInfo()
	if r != nil {
		this.Data["json"] = *r
	}
	this.ServeJson()
}