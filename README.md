# go_blog
这个一个简单的博客，仿照我之前用django写的博客。
现在用go重新写一遍。没有django自带的后台管理系统，只能自己写了。
没有django丰富的三方件，写起来真的蛋疼。

运行说明
克隆代码之后，修改models/blog.go里面的
orm.RegisterDataBase("default", "mysql", "用户名:密码@tcp(数据库ip:端口号)/gosql?charset=utf8")
我使用的是mysql数据库。
修改main.go里面的
beego.SetStaticPath("/static", 代码路径+"\\go_blog\\static")
之后go run main.go即可运行
打开
http://127.0.0.1:8080
进入主页面
http://127.0.0.1:8080/admin
进入管理页面
分页功能待开发
管理页面ajax异步搜索待开发