package models

type LoginUser struct{
	Username string    `form:"username"`
	Password string    `form:"password"`
}

type RegisterUser struct {
	Username   string    `form:"username"`
	Password   string    `form:"password"`
	ConfirmPwd string    `form:"confirmPwd"`
}

type BlogInfo struct {
	Id       int        `form:"id"`
	Title    string     `form:"title"`
	Content  string     `form:"content"`
	Category []int     `form:"categorys"`
}

type CategoryInfo struct {
	Id          int         `form:"id"`
	Name        string      `form:"name"`
	Description string      `form:"description"`
}
