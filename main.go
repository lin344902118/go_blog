package main

import (
	"github.com/astaxie/beego"
	_ "go_blog/models"
	_ "go_blog/routers"
)

func main() {
	beego.SetStaticPath("/static", "D:\\gopath\\src\\go_blog\\static")
	beego.SetLogger("file", `{"filename": "test.log"}`)
	beego.BeeLogger.DelLogger("console")
	beego.Run()
}

