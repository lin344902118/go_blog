package utils

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go_blog/models"
)

func RegisterUser(username, password string) (error) {
	o := orm.NewOrm()
	user := models.User{Username:username,Password:Md5Encrypted(password)}
	_, err := o.Insert(&user)
	if err != nil {
		beego.Warn("register user failed, err", err)
		return errors.New("register error")
	} else {
		return nil
	}
}

func GetUser(fieldName string, filedValue interface{}) (models.User, error) {
	o := orm.NewOrm()
	var user models.User
	err := o.QueryTable(user).Filter(fieldName, filedValue).RelatedSel().One(&user)
	return user, err
}
