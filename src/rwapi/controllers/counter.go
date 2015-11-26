package controllers
import (
	"github.com/astaxie/beego"
	"rwapi/code"
	"rwapi/common"
)

type CounterController struct {
	beego.Controller
}

//@router /counter/:id [get]
func (this *CounterController)GetCounter(){
	params := this.Ctx.Input.Params
	id := params[":id"]
	model := common.ConvertQueryToCmd(this.Ctx.Input)
	s := code.Exec(id,model)
	if s != nil {
		this.Data["json"] = *s
	}else {
		this.Ctx.Output.Status = 502
	}
	this.ServeJson()
}