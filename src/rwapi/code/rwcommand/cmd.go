package rwcommand
import (
	"rwapi/models"
	"rwapi/code/rwcache"
)

type Cmd struct {
	Message string
	ReturnMessage string
}

//命令处理
func(this *Cmd) Process(appkey string)( []models.AppDataModel){

	var list = rwcache.GetDatas(appkey)

	return list
}

