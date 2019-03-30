package routers

import (
	"go_blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{}, "get:Index")
    beego.Router("/logout", &controllers.MainController{}, "get:Logout")
    beego.Router("/about", &controllers.MainController{}, "get:About")
    beego.Router("/search", &controllers.MainController{}, "post:Search")
    beego.Router("/article", &controllers.MainController{}, "get,post:Article")
    beego.Router("/admin", &controllers.MainController{}, "get:Admin")
    beego.Router("/login", &controllers.MainController{}, "get,post:Login")
    beego.Router("/register", &controllers.MainController{}, "get,post:Register")
    beego.Router("/admin/blog", &controllers.MainController{}, "get:Blog")
    beego.Router("/admin/category", &controllers.MainController{}, "get:Category")
}
