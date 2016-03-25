package controllers
import (
	"github.com/astaxie/beego"
	"rwapi/code"
)

type HeartBeatController struct {
	beego.Controller
}

//@router /info/:id [get]
func (this *HeartBeatController)GetHeartBeat(){
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