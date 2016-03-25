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

//连接Appkey状态
//@router /keys [get]
func (this *StatusControllers)GetAllKey()  {
	r := code.AppKeyInfo()
	if r != nil {
		this.Data["json"] = *r
	}
	this.ServeJson()
}