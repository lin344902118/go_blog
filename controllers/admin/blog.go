package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"go_blog/models"
	"go_blog/utils"
	"time"
)

func (c *AdminController) Blog() {
	getUserOrRedirectLogin(c)
	RenderLayout(c, "blog")
	getAndRenderBlogs(c)
}

func (c *AdminController) EditBlog() {
	getUserOrRedirectLogin(c)
	RenderLayout(c, "blog")
	if c.Ctx.Input.Method() == "GET" {
		getEditBlog(c)
	} else if c.Ctx.Input.Method() == "POST" {
		postEditBlog(c)
	} else {
		c.Abort("405")
	}
}

func (c *AdminController) BlogDetail() {
	getUserOrRedirectLogin(c)
	RenderLayout(c, "blog")
	getBlogDetail(c)
}

func (c *AdminController) DeleteBlog() {
	var blogId int
	var ret = 1
	var message = ""
	_, err := GetUserBySession(c)
	if err != nil {
		message = utils.USER_NOT_LOGIN
	} else {
		if err := c.Ctx.Input.Bind(&blogId, "id"); err != nil {
			message = utils.ID_NO_FOUND
		} else {
			if err = utils.DeleteBlog(blogId); err != nil {
				message = utils.DELETE_BLOG_ERROR
			} else {
				ret = 0
				message = "删除成功"
			}
		}
	}
	c.Data["json"] = map[string]interface{}{"ret":ret,"message":message}
	c.ServeJSON()
}


func getAndRenderBlogs(c *AdminController) {
	if blogs, err := utils.GetAllBlogs(); err != nil {
		c.Data["error"] = utils.GET_BLOG_DATA_ERROR
	} else {
		c.Data["blogs"] = blogs
	}
	c.TplName = "blog.html"
}

func postEditBlog(c *AdminController) {
	var blogInfo BlogInfo
	if err := c.ParseForm(&blogInfo); err != nil {
		c.Data["error"] = utils.PARSE_BLOG_DATA_ERROR
	} else {
		categorys := getCategorysById(blogInfo, c)
		createOrUpdateBlog(c, blogInfo, categorys)
	}
	c.TplName = "blog.html"
}

func getCategorysById(blogInfo BlogInfo, c *AdminController) []*models.Category {
	var categorys []*models.Category
	for _, id := range blogInfo.Category {
		category, err := utils.GetCategory("id", id)
		if err != nil {
			c.Data["error"] = utils.ID_NO_FOUND
		} else {
			fmt.Println("category", category)
			categorys = append(categorys, category)
		}
	}
	return categorys
}

func createOrUpdateBlog(c *AdminController, blogInfo BlogInfo, categorys []*models.Category) {
	userId := c.GetSession("userId")
	newBlog := models.Blog{Title: blogInfo.Title,PublicTime:time.Now(),
		Content: blogInfo.Content, Author: &models.User{Id: userId.(int)}}
	if _, err := utils.GetBlog("Id", blogInfo.Id); err != nil {
		fmt.Println("create blog")
		createBlog(newBlog, c, categorys)
	} else {
		fmt.Println("update blog")
		newBlog.Id = blogInfo.Id
		updateBlog(newBlog, c, categorys)
	}
}

func createBlog(newBlog models.Blog, c *AdminController, categorys []*models.Category) {
	// blog not exist, insert
	if id, err := utils.CreateBlogWithCategorys(newBlog, categorys); err != nil {
		beego.Warn("create blog error.err:", err)
		c.Data["error"] = utils.CREATE_BLOG_ERROR
	} else {
		c.Redirect(fmt.Sprintf("/admin/blog/detail?id=%d", id), 302)
	}
}

func updateBlog(newBlog models.Blog,  c *AdminController, categorys []*models.Category) {
	// blog exist, update
	if err := utils.UpdateBlogWithCategory(newBlog, categorys); err != nil {
		beego.Warn("update blog error.err:", err)
		c.Data["error"] = utils.UPDATE_BLOG_ERROR
	} else {
		c.Redirect(fmt.Sprintf("/admin/blog/detail?id=%d", newBlog.Id), 302)
	}
}

func getEditBlog(c *AdminController) {
	var blogId int
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Script"] = "tinymceScript.html"
	c.Data["edit"] = "true"
	c.Layout = "admin.html"
	getAndRenderCategorys(c)
	// overwrite TplName
	c.TplName = "editBlog.html"
	if err := c.Ctx.Input.Bind(&blogId, "id"); err == nil && blogId != 0 {
		getAndRenderBlog(blogId, c)
	}
}

func getAndRenderBlog(blogId int, c *AdminController) {
	blog, err := utils.GetBlog("Id", blogId)
	if err != nil {
		c.Data["error"] = utils.ID_ERROR
	} else {
		c.Data["blog"] = blog
	}
}

func getBlogDetail(c *AdminController) {
	var blogId int
	if err := c.Ctx.Input.Bind(&blogId, "id"); err != nil || blogId == 0  {
		c.Data["error"] = utils.ID_NO_FOUND
	} else {
		blog, err := utils.GetBlogWithCategorys("Id", blogId)
		if err != nil {
			c.Data["error"] = utils.GET_BLOG_DATA_ERROR
		} else {
			c.Data["blog"] = blog
		}
	}
	c.Data["edit"] = "true"
	c.TplName = "blogDetail.html"
}
