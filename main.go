package main

import (
	"github.com/astaxie/beego"
	_ "go_blog/models"
	_ "go_blog/routers"
	"go_blog/utils"
)

func main() {
	beego.SetStaticPath("/static", "D:\\gopath\\src\\go_blog\\static")
	beego.SetLogger("file", `{"filename": "test.log"}`)
	beego.BeeLogger.DelLogger("console")
	beego.AddFuncMap("translate", utils.Translates)
	beego.AddFuncMap("formatTime", utils.FormatTime)
	beego.Run()
}

