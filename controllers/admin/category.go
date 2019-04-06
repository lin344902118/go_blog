package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go_blog/models"
	"go_blog/utils"
)

func (c *AdminController) Category() {
	getUserOrRedirectLogin(c)
	RenderLayout(c, "category")
	getAndRenderCategorys(c)
}

func (c *AdminController) EditCategory() {
	getUserOrRedirectLogin(c)
	RenderLayout(c, "category")
	if c.Ctx.Input.Method() == "GET" {
		getEditCategory(c)
	} else if c.Ctx.Input.Method() == "POST" {
		postEditCategory(c)
	} else {
		c.Abort("405")
	}
}

func (c *AdminController) CategoryDetail() {
	getUserOrRedirectLogin(c)
	RenderLayout(c, "category")
	getCategoryDetail(c)
}

func (c *AdminController) DeleteCategory() {
	var categoryId int
	var ret = 1
	var message = ""
	_, err := GetUserBySession(c)
	if err != nil {
		message = utils.USER_NOT_LOGIN
	} else {
		if err := c.Ctx.Input.Bind(&categoryId, "id"); err != nil {
			message = utils.ID_NO_FOUND
		} else {
			if err = utils.DeleteCategory(categoryId); err != nil {
				message = utils.DELETE_CATEGORY_ERROR
			} else {
				ret = 0
				message = "删除成功"
			}
		}
	}
	c.Data["json"] = map[string]interface{}{"ret":ret,"message":message}
	c.ServeJSON()
}

func getAndRenderCategorys(c *AdminController) {
	categorys, err := utils.GetAllCategorys()
	if err != nil {
		c.Data["error"] = utils.GET_CATEGORY_DATA_ERROR
	} else {
		c.Data["categorys"] = categorys
	}
	c.TplName = "category.html"
}

func postEditCategory(c *AdminController) {
	c.TplName = "category.html"
	var categoryInfo CategoryInfo
	if err := c.ParseForm(&categoryInfo); err != nil {
		c.Data["error"] = utils.PARSE_CATEGORY_DATA_ERROR
	} else {
		createOrUpdateCategory(categoryInfo, c)
	}
}

func createOrUpdateCategory(categoryInfo CategoryInfo, c *AdminController) {
	newCategory := models.Category{Id: categoryInfo.Id,
		Name: categoryInfo.Name, Description: categoryInfo.Description}
	if _, err := utils.GetCategory("Id", categoryInfo.Id); err != nil {
		insertCategory(newCategory, c)
	} else {
		newCategory.Id = categoryInfo.Id
		updateCategory(newCategory, c)
	}
}

func updateCategory(newCategory models.Category, c *AdminController) {
	o := orm.NewOrm()
	_, err := o.Update(&newCategory)
	if err != nil {
		beego.Warn(fmt.Sprintf("update category error.err:%s", err))
		c.Data["error"] = utils.UPDATE_CATEGORY_ERROR
	} else {
		c.Redirect(fmt.Sprintf("/admin/category/detail?id=%d", newCategory.Id), 302)
	}
}

func insertCategory(category models.Category, c *AdminController) {
	o := orm.NewOrm()
	id, err := o.Insert(&category)
	if err != nil {
		beego.Warn(fmt.Sprintf("insert category error.err:%s", err))
		c.Data["error"] = utils.CREATE_CATEGORY_ERROR
	} else {
		c.Redirect(fmt.Sprintf("/admin/category/detail?id=%d", id), 302)
	}
}

func getEditCategory(c *AdminController) {
	var categoryId int
	c.Data["edit"] = "true"
	c.Layout = "admin.html"
	c.TplName = "editCategory.html"
	if err := c.Ctx.Input.Bind(&categoryId, "id"); err == nil && categoryId != 0 {
		getAndRenderCategory(categoryId, c)
	}
}

func getAndRenderCategory(categoryId int, c *AdminController) {
	category, err := utils.GetCategory("Id", categoryId)
	if err != nil {
		c.Data["error"] = utils.ID_ERROR
	} else {
		c.Data["category"] = category
	}
}

func getCategoryDetail(c *AdminController) {
	var categoryId int
	if err := c.Ctx.Input.Bind(&categoryId, "id"); err != nil || categoryId == 0 {
		c.Data["error"] = utils.ID_NO_FOUND
	} else {
		category, _ := utils.GetCategory("Id", categoryId)
		c.Data["category"] = category
	}
	c.TplName = "categoryDetail.html"
}
