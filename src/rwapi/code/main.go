package code

import (
	"rwapi/code/rwnetwork"
	"rwapi/code/rwcommand"
	"rwapi/code/rwconnection"
	"rwapi/models"
	"rwapi/code/rwcache"
)

func BeginRun() {
	rwnetwork.Run()
}


func SystemInfo(id string) *models.ReturnModel {

	model := new(models.ReturnModel)
	model.Data = ""
	model.Message = ""
	return model
}

func StatusInfo() *models.ReturnModel {
	d := rwconnection.GetAllAddr()
	model := new(models.ReturnModel)
	model.Data = d
	return model
}

func AppKeyInfo()*models.ReturnModel  {
	d:=rwcache.GetAllKey()
	model := new(models.ReturnModel)
	model.Data = d
	return model
}

func Exec(appkey string) *models.ReturnModel {
	var cmd = new(rwcommand.Cmd)

	var d = cmd.Process(appkey)

	model := new(models.ReturnModel)
	model.Data = d
	model.Message = ""
	return model
}