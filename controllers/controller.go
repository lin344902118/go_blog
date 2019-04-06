package controllers

import (
	"github.com/astaxie/beego"
	"go_blog/utils"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	c.Layout = "layout.html"
	blogs, err := utils.GetAllBlogs()
	if err != nil {
		c.Data["error"] = utils.GET_BLOG_DATA_ERROR
	} else {
		c.Data["blogs"] = blogs
	}
	c.TplName = "home.html"
}

func (c *MainController) About() {
	c.Layout = "layout.html"
	c.TplName = "about.html"
}

func (c *MainController) Search() {
}

func (c *MainController) Article() {
	c.Layout = "layout.html"
	c.TplName = "home.html"
}

