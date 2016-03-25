package controllers
import (
	"github.com/astaxie/beego"
	"rwapi/code"
)

type CounterController struct {
	beego.Controller
}

//@router /counter/:id [get]
func (this *CounterController)GetCounter(){
	params := this.Ctx.Input.Params
	id := params[":id"]
	s := code.Exec(id)
	if s != nil {
		this.Data["json"] = *s
	}else {
		this.Ctx.Output.Status = 502
	}
	this.ServeJson()
}