package admin

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"go_blog/models"
	"go_blog/utils"
)

func GetUser(c *AdminController) (models.User, error) {
	user, err := GetUserBySession(c)
	if err != nil {
		beego.Warn(fmt.Sprintf("get user failed, error: %s", err))
		return user, errors.New("user not login")
	} else {
		return user, nil
	}
}

func RenderLayout(c *AdminController, current string) {
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

func GetUserBySession(c *AdminController) (models.User, error) {
	var user models.User
	userId := c.GetSession("userId")
	if userId == nil {
		return user, errors.New("user Id not found")
	} else {
		user, err := utils.GetUser("Id", userId.(int))
		return user, err
	}
}

func RedirectToLogin(c *AdminController) {
	c.Redirect("/login", 302)
}

func getUserOrRedirectLogin(c *AdminController) {
	if _, err := GetUser(c); err != nil {
		RedirectToLogin(c)
	}
}

