package controllers

import (
	"github.com/astaxie/beego"
	"rwapi/code"
	"rwapi/code/rwcache"
	"wechatjob/model"
)

type CommandController struct {
	beego.Controller
}


//@router /exec/add [post]
func (this *CommandController)AddMsg()  {
	var msg = this.Ctx.Input.RequestBody
	rwcache.AddMsg(string(msg))
	this.Data["json"] = new(model.ReturnModel)
	this.ServeJson()
}


//@router /exec/:id [get]
func (this *CommandController)Exec() {
	var appkey = this.Ctx.Input.Param(":id")
	r := code.Exec(appkey)
	if r != nil {
		this.Data["json"] = *r
	}else {
		this.Data["json"] = nil
	}
	this.ServeJson()
}