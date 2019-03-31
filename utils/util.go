package utils

import (
	"github.com/astaxie/beego/orm"
	"go_blog/models"
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

func GetBlog(fieldName string, fieldValue interface{}) (models.Blog, error){
	o := orm.NewOrm()
	var blog models.Blog
	err := o.QueryTable("blog").Filter(fieldName, fieldValue).One(&blog)
	return blog, err
}

func GetCategory(fieldName string, fieldValue interface{}) (models.Category, error){
	o := orm.NewOrm()
	var category models.Category
	err := o.QueryTable("blog").Filter(fieldName, fieldValue).One(&category)
	return category, err
}