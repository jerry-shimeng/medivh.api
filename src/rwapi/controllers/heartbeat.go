package controllers
import (
	"github.com/astaxie/beego"
	"rwapi/common"
	"rwapi/code"
)

type HeartBeatController struct {
	beego.Controller
}

//@router /info/:id [get]
func (this *HeartBeatController)GetHeartBeat(){
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