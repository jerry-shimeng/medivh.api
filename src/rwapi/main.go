package main

import (
	_ "rwapi/docs"
	_ "rwapi/routers"
	"rwapi/code"
	"github.com/astaxie/beego"
)

func main() {

	//启动监视模块
	go code.BeginRun()

	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()


}
