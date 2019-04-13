package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"go_blog/models"
	"go_blog/utils"
	"time"
)

func (c *AdminController) GetBlog() {
	director := GetBlogDirector(c, &GetBlog{})
	director.getModel()
}

func (c *AdminController) EditBlog() {
	director := GetBlogDirector(c, &EditBlog{})
	director.getModel()
}

func (c *AdminController) PostBlog() {
	director := GetBlogDirector(c, &PostBlog{})
	director.getModel()
}

func (c *AdminController) BlogDetail() {
	director := GetBlogDirector(c, &BlogDetail{})
	director.getModel()
}

func (c *AdminController) DeleteBlog() {
	DeleteRecordAndReturnJson(c, utils.DeleteBlog, utils.DELETE_BLOG_ERROR)
}

type GetBlog struct {
	Admin
}

func (self *GetBlog) RenderData(c *AdminController){
	if blogs, err := utils.GetAllBlogs(); err != nil {
		c.Data["error"] = utils.GET_BLOG_DATA_ERROR
	} else {
		c.Data["blogs"] = blogs
	}
	c.TplName = "blog.html"
}

type EditBlog struct {
	Admin
}

func (self *EditBlog) RenderData(c *AdminController) {
	var blogId int
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Script"] = "tinymceScript.html"
	c.Data["edit"] = "true"
	c.Layout = "admin.html"
	categorys := GetCategory{}
	categorys.RenderData(c)
	// overwrite TplName
	c.TplName = "editBlog.html"
	if err := c.Ctx.Input.Bind(&blogId, "id"); err == nil && blogId != 0 {
		getAndRenderBlog(blogId, c)
	}
}

type PostBlog struct {
	Admin
}

func (self *PostBlog) RenderData(c *AdminController) {
	var blogInfo models.BlogInfo
	if err := c.ParseForm(&blogInfo); err != nil {
		c.Data["error"] = utils.PARSE_BLOG_DATA_ERROR
	} else {
		categorys := getCategorysById(blogInfo, c)
		createOrUpdateBlog(c, blogInfo, categorys)
	}
	c.TplName = "blog.html"
}

type BlogDetail struct {
	Admin
}

func (self *BlogDetail) RenderData(c *AdminController)  {
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

func getCategorysById(blogInfo models.BlogInfo, c *AdminController) []*models.Category {
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

func createOrUpdateBlog(c *AdminController, blogInfo models.BlogInfo, categorys []*models.Category) {
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

func getAndRenderBlog(blogId int, c *AdminController) {
	blog, err := utils.GetBlog("Id", blogId)
	if err != nil {
		c.Data["error"] = utils.ID_ERROR
	} else {
		c.Data["blog"] = blog
	}
}