# go_blog
这个一个简单的博客，仿照我之前用django写的博客。
现在用go重新写一遍。没有django自带的后台管理系统，只能自己写了。
没有django丰富的三方件，写起来真的蛋疼。

运行说明
克隆代码之后，修改models/blog.go里面的
orm.RegisterDataBase("default", "mysql", "用户名:密码@tcp(数据库ip:端口号)/gosql?charset=utf8")
我使用的是mysql数据库。
之后go run main.go即可运行
打开http://127.0.0.1:8080就能看到了
后续有时间会完善搜索和分页功能。