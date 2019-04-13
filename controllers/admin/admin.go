package admin

import (
	"github.com/astaxie/beego"
	"go_blog/utils"
)

type AdminController struct {
	beego.Controller
}

type AdminDirector struct {
	controller *AdminController
	modelBuilder AdminModel
	current string
}

func (self *AdminDirector) getModel() {
	self.modelBuilder.GetUserOrRedirectLogin(self.controller)
	self.modelBuilder.RenderLayout(self.controller, self.current)
	self.modelBuilder.RenderData(self.controller)
}

func GetBlogDirector(c *AdminController, builder AdminModel) AdminDirector {
	return AdminDirector{c, builder, "blog"}
}

func GetCategoryDirector(c *AdminController, builder AdminModel) AdminDirector {
	return AdminDirector{c, builder, "category"}
}

type AdminModel interface {
	GetUserOrRedirectLogin(c *AdminController)
	RenderLayout(c *AdminController, current string)
	RenderData(c *AdminController)
}

type Admin struct {
}

func (self *Admin)GetUserOrRedirectLogin(c *AdminController) {
	if _, err := GetUser(c); err != nil {
		c.Redirect("/login", 302)
	}
}

func (self *Admin) RenderLayout(c *AdminController, current string) {
	c.Data["current"] = current
	tables := []map[string]string{}
	tableNames := utils.GetAllTableNames()
	for _, table := range tableNames {
		if table == current {
			tables = append(tables, map[string]string{"name": current, "active": "true"})
		} else {
			tables = append(tables, map[string]string{"name": table, "active": "false"})
		}
	}
	c.Data["tables"] = tables
	c.Layout = "admin.html"
}

func DeleteRecordAndReturnJson(c *AdminController, DeleteFunction func(int)(error), errMsg string) {
	var recordId int
	var ret = 1
	var message = ""
	_, err := GetUserBySession(c)
	if err != nil {
		message = utils.USER_NOT_LOGIN
	} else {
		if err := c.Ctx.Input.Bind(&recordId, "id"); err != nil {
			message = utils.ID_NO_FOUND
		} else {
			if err = DeleteFunction(recordId); err != nil {
				message = errMsg
			} else {
				ret = 0
				message = "删除成功"
			}
		}
	}
	c.Data["json"] = map[string]interface{}{"ret":ret,"message":message}
	c.ServeJSON()
}
