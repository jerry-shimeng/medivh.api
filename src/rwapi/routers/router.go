// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"rwapi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/s",
			beego.NSInclude(
				&controllers.SystemController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/status",
			beego.NSInclude(
				&controllers.StatusControllers{},
			),
		),
		beego.NSNamespace("/c",
			beego.NSInclude(
				&controllers.CounterController{},
			),
		),
		beego.NSNamespace("/cmd",
			beego.NSInclude(
				&controllers.CommandController{},
			),
		),
	)
	beego.AddNamespace(ns)
//	beego.Include(&controllers.SystemController{})
//	beego.Include(&controllers.UserController{})
//	beego.Include(&controllers.StatusControllers{})
}
