package main

import (
	_ "rwapi/routers"
	"github.com/astaxie/beego"
	"rwapi/code"
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
