package utils

import (
	"go_blog/models"
	"github.com/astaxie/beego/orm"
	"math/rand"
)

const letters  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GetRandomString(length int) string{
	randStr := make([]byte, length)
	for i := range randStr {
		randStr[i] = letters[rand.Intn(length)]
	}
	return string(randStr)
}

func GetAllBlogs() ([]models.Blog, error){
	o := orm.NewOrm()
	var blogs []models.Blog
	_, err := o.QueryTable("blog").All(&blogs)
	if err != nil {
		return nil, err
	} else {
		return blogs, nil
	}
}

func GetAllCategorys() ([]models.Category, error){
	o := orm.NewOrm()
	var categorys []models.Category
	_, err := o.QueryTable("category").All(&categorys)
	if err != nil {
		return nil, err
	} else {
		return categorys, nil
	}
}
