package controllers
import (
	"github.com/astaxie/beego"
	"rwapi/models"
	"encoding/json"
	"rwapi/code"
)

type CommandController struct {
	beego.Controller
}


//@router /:id [post]
func (this *CommandController)Exec(){
	var ob models.CmdModel
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	if err != nil {
		this.Data["json"] = models.ReturnModel{Code:-1,Message:err.Error()}
	}else {
		if ob.Module <=0 {
			this.Data["json"] = models.ReturnModel{Code:-1,Message:"the bad module"}
		}else {
			r := code.Exec(this.Ctx.Input.Param(":id"),&ob)
			if r != nil {
				this.Data["json"] = *r
			}else {
				this.Data["json"] = nil
			}
		}
	}

	this.ServeJson()
}