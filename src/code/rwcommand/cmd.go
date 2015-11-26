package rwcommand
import (
	"rwapi/models"
	"rwapi/code/rwconnection"
	"errors"
	"encoding/json"
)

type Cmd struct {
	Message string
	ReturnMessage string
}

//命令处理
func(this *Cmd) Process(id string ,cmd *models.CmdModel)( *interface{}, error){
	//找到连接对象
	model := rwconnection.FindConn(id)
	if model == nil {
		return nil,errors.New("not found connection")
	}
	//发送命令
	model.Send(cmd.ToCmdString())
	//接受返回消息
	s := model.Receive()
	var v = new(interface{})
	err := json.Unmarshal([]byte(*s),v)
	return v ,err
}

