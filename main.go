package main

import (
	_ "blog/models"
	_ "blog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static", "D:\\gopath\\src\\blog\\static")
	beego.SetLogger("file", `{"filename": "test.log"}`)
	beego.BeeLogger.DelLogger("console")
	beego.Run()
}

