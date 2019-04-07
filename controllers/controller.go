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
	blogs, err := utils.GetAllBlogsWithCategorys()
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
	q := c.Input().Get("q")
	if q == "" {
		c.Redirect("/index", 302)
	} else {
		blogs, err := utils.SearchBlog(q, 5)
		c.Layout = "layout.html"
		if err != nil {
			c.Data["error"] = utils.SEARCH_BLOG_ERROR
		} else {
			c.Data["blogs"] = blogs
		}
		c.TplName = "home.html"
	}

}

func (c *MainController) Article() {
	c.Layout = "layout.html"
	c.TplName = "home.html"
}

