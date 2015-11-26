package common
import (
	"rwapi/models"
	"github.com/astaxie/beego/context"
	"strconv"
	"fmt"
)

func ConvertQueryToCmd(input *context.BeegoInput) *models.CmdModel {

	model := new(models.CmdModel)
	model.Alias =input.Query("alias")
	model.Counter,_ = strconv.Atoi(input.Query("counter"))
	model.Level,_ = strconv.Atoi(input.Query("level"))
	model.Mark = input.Query("mark")
	model.Operate = input.Query("operate")
	model.Module,_ =strconv.Atoi(input.Query("module"))

	fmt.Printf("%#v",model)
	return model
}