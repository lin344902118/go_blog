package utils

import (
	"github.com/astaxie/beego/orm"
	"go_blog/models"
)

func GetAllTableNames() ([]string) {
	return []string{"blog", "category"}
}

func GetAllBlogs() ([]models.Blog, error) {
	o := orm.NewOrm()
	var blogs []models.Blog
	_, err := o.QueryTable("blog").All(&blogs)
	if err != nil {
		return nil, err
	} else {
		return blogs, nil
	}
}

func GetAllCategorys() ([]models.Category, error) {
	o := orm.NewOrm()
	var categorys []models.Category
	_, err := o.QueryTable("category").All(&categorys)
	if err != nil {
		return nil, err
	} else {
		return categorys, nil
	}
}

func GetBlog(fieldName string, fieldValue interface{}) (*models.Blog, error) {
	o := orm.NewOrm()
	var blog models.Blog
	err := o.QueryTable("blog").Filter(fieldName, fieldValue).One(&blog)
	return &blog, err
}

func GetCategory(fieldName string, fieldValue interface{}) (*models.Category, error) {
	o := orm.NewOrm()
	var category models.Category
	err := o.QueryTable("category").Filter(fieldName, fieldValue).One(&category)
	return &category, err
}

func DeleteBlog(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&models.Blog{Id:id})
	return err
}

func DeleteCategory(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&models.Category{Id:id})
	return err
}