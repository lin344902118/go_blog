package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go_blog/models"
	"go_blog/utils"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	c.Layout = "layout.html"
	blogs, err := utils.GetAllBlogs()
	if err != nil {
		c.Data["error"] = "获取博文数据失败"
		c.TplName = "error.html"
	} else {
		c.Data["blogs"] = blogs
		c.TplName = "home.html"
	}
}

func (c *MainController) Logout() {
	c.DelSession("userId")
	c.Ctx.Redirect(302, "/")
}

func (c *MainController) About() {
	c.Layout = "layout.html"
	c.TplName = "about.html"
}

func (c *MainController) Search() {

}

func (c *MainController) Article() {

}

func (c *MainController) Admin() {
	userId := c.GetSession("userId")
	if userId == nil {
		c.TplName = "login.html"
	} else {
		user, err := utils.GetUser("id", userId.(int))
		if err != nil {
			c.TplName = "login.html"
		} else {
			c.Data["username"] = user.Username
			c.Data["current"] = "blog"
			tables := []map[string]string{}
			tables = append(tables, map[string]string{"name": "blog", "active": "true"})
			tables = append(tables, map[string]string{"name": "category", "active": "false"})
			c.Data["tables"] = tables
			c.TplName = "admin.html"
		}
	}
}

func (c *MainController) Login() {
	c.TplName = "login.html"
	if c.Ctx.Input.Method() == "POST" {
		username := c.Input().Get("username")
		password := c.Input().Get("password")
		user, err := utils.GetUser("username", username)
		if err != nil {
			c.Data["errMsg"] = "用户不存在"
		} else {
			if user.Id != 0 {
				encrypt := utils.Md5Encrypted(password)
				if encrypt == user.Password {
					// login successfully
					c.SetSession("userId", user.Id)
					c.Redirect("/admin", 301)
				} else {
					c.Data["errMsg"] = "用户名密码不正确"
				}
			}
		}
	}
}

func (c *MainController) Register() {
	c.TplName = "register.html"
	if c.Ctx.Input.Method() == "POST" {
		username := c.Input().Get("username")
		password := c.Input().Get("password")
		confirmPwd := c.Input().Get("confirmPwd")
		if password != confirmPwd {
			c.Data["errMsg"] = "两次密码不一致"
		} else {
			user, err := utils.GetUser("username", username)
			if err == nil {
				c.Data["errMsg"] = "用户已存在"
			} else {
				err := utils.RegisterUser(username, password)
				if err != nil {
					c.Data["errMsg"] = "注册失败"
				} else {
					// register successfully
					c.SetSession("userId", user.Id)
					c.Redirect("/admin", 301)
				}
			}
		}
	}
}

func (c *MainController) Blog() {
	userId := c.GetSession("userId")
	if userId == nil {
		c.TplName = "login.html"
	} else {
		user, err := utils.GetUser("id", userId.(int))
		if err != nil {
			c.TplName = "login.html"
		} else {
			c.Data["username"] = user.Username
			c.Data["current"] = "blog"
			tables := []map[string]string{}
			tables = append(tables, map[string]string{"name": "blog", "active": "true"})
			tables = append(tables, map[string]string{"name": "category", "active": "false"})
			c.Data["tables"] = tables
			c.Layout = "admin.html"
			blogs, err := utils.GetAllBlogs()
			if err != nil {
				c.Data["error"] = "获取博文数据失败"
				c.TplName = "error.html"
			} else {
				c.Data["blogs"] = blogs
				c.TplName = "article.html"
			}
		}
	}
}

func (c *MainController) Category() {
	userId := c.GetSession("userId")
	if userId == nil {
		c.TplName = "login.html"
	} else {
		user, err := utils.GetUser("id", userId.(int))
		if err != nil {
			c.TplName = "login.html"
		}  else {
			c.Data["username"] = user.Username
			c.Data["current"] = "category"
			tables := []map[string]string{}
			tables = append(tables, map[string]string{"name": "blog", "active": "false"})
			tables = append(tables, map[string]string{"name": "category", "active": "true"})
			c.Data["tables"] = tables
			c.Layout = "admin.html"
			categorys, err := utils.GetAllCategorys()
			if err != nil {
				c.Data["error"] = "获取类别数据失败"
				c.TplName = "error.html"
			} else {
				c.Data["categorys"] = categorys
				c.TplName = "category.html"
			}
		}
	}
}

func (c *MainController) EditBlog() {
	userId := c.GetSession("userId")
	if userId == nil {
		c.TplName = "login.html"
	} else {
		user, err := utils.GetUser("id", userId.(int))
		if err != nil {
			c.TplName = "login.html"
		}  else {
			c.Data["username"] = user.Username
			c.Data["current"] = "blog"
			tables := []map[string]string{}
			tables = append(tables, map[string]string{"name": "blog", "active": "true"})
			tables = append(tables, map[string]string{"name": "category", "active": "false"})
			c.Data["tables"] = tables
			categorys, _ := utils.GetAllCategorys()
			c.Data["categorys"] = categorys
			c.Layout = "admin.html"
			if c.Ctx.Input.Method() == "GET"{
				blogId := c.Input().Get("id")
				c.LayoutSections = make(map[string]string)
				c.LayoutSections["Script"] = "tinymceScript.html"
				c.Data["edit"] = "true"
				if blogId == "" {
					c.TplName = "editBlog.html"
				} else {
					id, err := strconv.Atoi(blogId)
					if err != nil {
						c.Data["error"] = "输入的id有误"
					} else {
						blog, err := utils.GetBlog("Id", id)
						if err != nil {
							c.Data["error"] = "id不存在"
						} else {
							c.Data["blog"] = blog
						}
					}
				}
			}
		}
	}
}

func (c *MainController) EditCategory() {
	userId := c.GetSession("userId")
	if userId == nil {
		c.TplName = "login.html"
	} else {
		user, err := utils.GetUser("id", userId.(int))
		if err != nil {
			c.TplName = "login.html"
		}  else {
			c.Data["username"] = user.Username
			c.Data["current"] = "category"
			tables := []map[string]string{}
			tables = append(tables, map[string]string{"name": "blog", "active": "false"})
			tables = append(tables, map[string]string{"name": "category", "active": "true"})
			c.Data["tables"] = tables
			c.Layout = "admin.html"
			if c.Ctx.Input.Method() == "GET"{
				categoryId := c.Input().Get("id")
				c.Data["edit"] = "true"
				if categoryId == "" {
					c.TplName = "editCategory.html"
				} else {
					id, err := strconv.Atoi(categoryId)
					if err != nil {
						c.Data["error"] = "输入的id有误"
					} else {
						category, err := utils.GetCategory("Id", id)
						if err != nil {
							c.Data["error"] = "id不存在"
						} else {
							c.Data["category"] = category
						}
					}
				}
			} else {
				categoryName := c.Input().Get("name")
				categoryDescription := c.Input().Get("description")
				category := models.Category{Name:categoryName,Description:categoryDescription}
				fmt.Println("category", category)
				o := orm.NewOrm()
				id, err := o.Insert(&category)
				if err != nil {
					c.Data["error"] = "创建分类失败"
				} else {
					fmt.Println("id", id)
					c.Redirect("/admin/category", 301)
				}
			}
		}
	}
}