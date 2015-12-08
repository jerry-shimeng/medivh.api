package code
import (
	"rwapi/code/rwnetwork"
	"rwapi/code/rwcommand"
	"rwapi/code/rwconnection"
	"rwapi/models"
	"fmt"
)

func BeginRun()  {
	rwnetwork.Run()
}


func SystemInfo(id string)*models.ReturnModel{
	var cmd = new(rwcommand.Cmd)
	//id为appkey，需要将其转换为objectid
	id = rwconnection.GetConnDict(id)

	var cmdModel = new(models.CmdModel)
	cmdModel.Module=2
	cmdModel.Alias ="sys"
	cmdModel.Counter=0
	model := new(models.ReturnModel)
	d,err := cmd.Process(id ,cmdModel)

	if err!=nil {
		model.Code = -1
	}
	model.Data = d
	model.Message = fmt.Sprint(err)
	return model
}

func StatusInfo()*models.ReturnModel{
	d := rwconnection.GetAllAddr()
	model := new(models.ReturnModel)
	model.Data = d
	return model
}

func Exec( id string ,cmdModel *models.CmdModel)*models.ReturnModel{
	//fmt.Printf("\n[XX] %v\n",*cmdModel)
	var cmd = new(rwcommand.Cmd)

	//id为appkey，需要将其转换为objectid
	id = rwconnection.GetConnDict(id)

	d,err := cmd.Process(id ,cmdModel)
	model := new(models.ReturnModel)
	if err!=nil {
		model.Code = -1
	}
	model.Data = d
	model.Message = fmt.Sprint(err)
	return model
}