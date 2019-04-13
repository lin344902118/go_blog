package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"go_blog/models"
	"go_blog/utils"
)

func (c *AdminController) GetCategory() {
	director := GetCategoryDirector(c, &GetCategory{})
	director.getModel()
}

func (c *AdminController) EditCategory() {
	director := GetCategoryDirector(c, &EditCategory{})
	director.getModel()
}

func (c *AdminController) PostCategory() {
	director := GetCategoryDirector(c, &PostCategory{})
	director.getModel()
}

func (c *AdminController) CategoryDetail() {
	director := GetCategoryDirector(c, &CategoryDetail{})
	director.getModel()
}

func (c *AdminController) DeleteCategory() {
	DeleteRecordAndReturnJson(c, utils.DeleteCategory, utils.DELETE_CATEGORY_ERROR)
}

type GetCategory struct {
	Admin
}

func (self *GetCategory) RenderData(c *AdminController) {
	categorys, err := utils.GetAllCategorys()
	if err != nil {
		c.Data["error"] = utils.GET_CATEGORY_DATA_ERROR
	} else {
		c.Data["categorys"] = categorys
	}
	c.TplName = "category.html"
}

type EditCategory struct {
	Admin
}

func (self *EditCategory) RenderData(c *AdminController) {
	var categoryId int
	c.Data["edit"] = "true"
	c.Layout = "admin.html"
	c.TplName = "editCategory.html"
	if err := c.Ctx.Input.Bind(&categoryId, "id"); err == nil && categoryId != 0 {
		getAndRenderCategory(categoryId, c)
	}
}

type PostCategory struct {
	Admin
}

func (self *PostCategory) RenderData(c *AdminController) {
	c.TplName = "category.html"
	var categoryInfo models.CategoryInfo
	if err := c.ParseForm(&categoryInfo); err != nil {
		c.Data["error"] = utils.PARSE_CATEGORY_DATA_ERROR
	} else {
		createOrUpdateCategory(categoryInfo, c)
	}
}

type CategoryDetail struct {
	Admin
}

func (self *CategoryDetail) RenderData(c *AdminController) {
	var categoryId int
	if err := c.Ctx.Input.Bind(&categoryId, "id"); err != nil || categoryId == 0 {
		c.Data["error"] = utils.ID_NO_FOUND
	} else {
		category, _ := utils.GetCategory("Id", categoryId)
		c.Data["category"] = category
	}
	c.Data["edit"] = "true"
	c.TplName = "categoryDetail.html"
}

func createOrUpdateCategory(categoryInfo models.CategoryInfo, c *AdminController) {
	newCategory := models.Category{Id: categoryInfo.Id,
		Name: categoryInfo.Name, Description: categoryInfo.Description}
	if _, err := utils.GetCategory("Id", categoryInfo.Id); err != nil {
		createCategory(newCategory, c)
	} else {
		newCategory.Id = categoryInfo.Id
		updateCategory(newCategory, c)
	}
}

func updateCategory(newCategory models.Category, c *AdminController) {
	err := utils.UpdateCategory(newCategory)
	if err != nil {
		beego.Warn(fmt.Sprintf("update category error.err:%s", err))
		c.Data["error"] = utils.UPDATE_CATEGORY_ERROR
	} else {
		c.Redirect(fmt.Sprintf("/admin/category/detail?id=%d", newCategory.Id), 302)
	}
}

func createCategory(category models.Category, c *AdminController) {
	id, err := utils.CreateCategory(category)
	if err != nil {
		beego.Warn(fmt.Sprintf("create category error.err:%s", err))
		c.Data["error"] = utils.CREATE_CATEGORY_ERROR
	} else {
		c.Redirect(fmt.Sprintf("/admin/category/detail?id=%d", id), 302)
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