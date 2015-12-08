package models
import "encoding/json"

type CmdModel struct {
	Mark,Alias,Operate string
	Level int
	Module int `json:"ModuleType"`
	Counter int `json:"CounterType"`
	Extend interface{}
}

func (this *CmdModel)ToJson()string{
	msg,_ := json.Marshal(this)
	return string(msg)
}

func (this *CmdModel)ToCmdString()string{

	return this.ToJson()+"\n"
}
