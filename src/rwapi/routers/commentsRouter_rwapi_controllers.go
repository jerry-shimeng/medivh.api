package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["rwapi/controllers:CommandController"] = append(beego.GlobalControllerRouter["rwapi/controllers:CommandController"],
		beego.ControllerComments{
			"Exec",
			`/:id`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["rwapi/controllers:CounterController"] = append(beego.GlobalControllerRouter["rwapi/controllers:CounterController"],
		beego.ControllerComments{
			"GetCounter",
			`/counter/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["rwapi/controllers:HeartBeatController"] = append(beego.GlobalControllerRouter["rwapi/controllers:HeartBeatController"],
		beego.ControllerComments{
			"GetHeartBeat",
			`/info/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["rwapi/controllers:StatusControllers"] = append(beego.GlobalControllerRouter["rwapi/controllers:StatusControllers"],
		beego.ControllerComments{
			"GetAll",
			`/all`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["rwapi/controllers:SystemController"] = append(beego.GlobalControllerRouter["rwapi/controllers:SystemController"],
		beego.ControllerComments{
			"GetInfo",
			`/info/:id`,
			[]string{"get"},
			nil})

}
