package utils

import (
	"go_blog/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)

func RegisterUser(username, password string) (error) {
	o := orm.NewOrm()
	user := models.User{Username:username,Password:Md5Encrypted(password)}
	_, err := o.Insert(&user)
	if err != nil {
		fmt.Println(err)
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
