package routers

import (
	"go_blog/controllers"
	"go_blog/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{}, "get:Index")
    beego.Router("/about", &controllers.MainController{}, "get:About")
    beego.Router("/search", &controllers.MainController{}, "get,post:Search")
    beego.Router("/article", &controllers.MainController{}, "get,post:Article")
    beego.Router("/admin", &admin.AdminController{}, "get:Blog")
    beego.Router("/login", &admin.AdminController{}, "get,post:Login")
    beego.Router("/logout", &admin.AdminController{}, "get:Logout")
    beego.Router("/register", &admin.AdminController{}, "get,post:Register")
    beego.Router("/admin/blog", &admin.AdminController{}, "get:Blog")
    beego.Router("/admin/category", &admin.AdminController{}, "get:Category")
    beego.Router("/admin/blog/edit", &admin.AdminController{}, "get,Post:EditBlog")
    beego.Router("/admin/category/edit", &admin.AdminController{}, "get,Post:EditCategory")
    beego.Router("/admin/blog/detail", &admin.AdminController{}, "get:BlogDetail")
    beego.Router("/admin/category/detail", &admin.AdminController{}, "get:CategoryDetail")
    beego.Router("/admin/blog/delete", &admin.AdminController{}, "post:DeleteBlog")
    beego.Router("/admin/category/delete", &admin.AdminController{}, "post:DeleteCategory")
}
