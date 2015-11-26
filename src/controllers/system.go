package controllers
import (
	"github.com/astaxie/beego"
	"rwapi/code"
)

type SystemController struct {
	beego.Controller
}

// @Title Get
// @Description get all
// @Success 200 {object} string
// @router /info/:id [get]
func (this *SystemController) GetInfo(){
	id := this.Ctx.Input.Param(":id")

	s:= code.SystemInfo(id)
	if s != nil {
		this.Data["json"] = *s
	}else {
		this.Ctx.Output.Status = 502
	}
	this.ServeJson()
}