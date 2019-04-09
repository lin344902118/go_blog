package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"go_blog/models"
	"go_blog/utils"
)

func (c *AdminController) Login() {
	c.TplName = "login.html"
	if c.Ctx.Input.Method() == "POST" {
		parseLoginInfo(c)
	} else {
		checkLoginAndRedirect(c)
	}
}

func (c *AdminController) Logout() {
	c.DelSession("userId")
	c.Redirect("/", 301)
}

func (c *AdminController) Register() {
	c.TplName = "register.html"
	if c.Ctx.Input.Method() == "POST" {
		parseRegisterInfo(c)
	}
}

func checkLoginAndRedirect(c *AdminController) {
	if _, err := GetUserBySession(c); err == nil {
		c.Redirect("/admin/blog", 302)
	}
}

func parseLoginInfo(c *AdminController) {
	var user models.LoginUser
	if err := c.ParseForm(&user); err != nil {
		c.Data["errMsg"] = utils.PARSE_USERNAME_PASSWORD_ERROR
	} else {
		checkLoginUserExists(user, c)
	}
}

func checkLoginUserExists(user models.LoginUser, c *AdminController) {
	u, err := utils.GetUser("Username", user.Username)
	if err != nil {
		c.Data["errMsg"] = utils.USER_NOT_EXIST
	} else {
		if u.Id != 0 {
			validatePassword(user, u, c)
		}
	}
}

func validatePassword(user models.LoginUser, u models.User, c *AdminController) {
	encrypt := utils.Md5Encrypted(user.Password)
	if encrypt == u.Password {
		// login successfully
		beego.Info(fmt.Sprintf("user:%s login successfully", u.Username))
		c.SetSession("userId", u.Id)
		c.Redirect("/admin/blog", 302)
	} else {
		c.Data["errMsg"] = utils.USERNAME_PASSWORD_ERROR
	}
}

func parseRegisterInfo(c *AdminController) {
	var rUser models.RegisterUser
	if err := c.ParseForm(&rUser); err != nil {
		c.Data["error"] = utils.PARSE_USERNAME_PASSWORD_ERROR
	} else {
		ValidateTwoPassword(rUser, c)
	}
}

func ValidateTwoPassword(rUser models.RegisterUser, c *AdminController) {
	if rUser.Password != rUser.ConfirmPwd {
		c.Data["errMsg"] = utils.TWO_PASSWORD_NOT_MATCH
	} else {
		checkRegisterUserExist(rUser, c)
	}
}

func checkRegisterUserExist(rUser models.RegisterUser, c *AdminController) {
	user, err := utils.GetUser("Username", rUser.Username)
	if err == nil {
		c.Data["errMsg"] = utils.USER_EXISTS
	} else {
		registerAndRedirect(rUser, c, user)
	}
}

func registerAndRedirect(rUser models.RegisterUser, c *AdminController, user models.User) {
	err := utils.RegisterUser(rUser.Username, rUser.Password)
	if err != nil {
		c.Data["errMsg"] = utils.REGISTER_FAILED
	} else {
		// register successfully
		beego.Info(fmt.Sprintf("user:%s register successfully", user.Username))
		c.SetSession("userId", user.Id)
		c.Redirect("/admin/blog", 302)
	}
}

