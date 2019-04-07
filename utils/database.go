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

func GetAllBlogsWithCategorys() ([]models.Blog, error) {
	blogs, err := GetAllBlogs()
	if err != nil {
		return blogs, err
	}
	for index, blog := range blogs {
		blogWithCategorys, _ := GetBlogWithCategorys("id", blog.Id)
		blogs[index] = *blogWithCategorys
	}
	return blogs, err
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

func GetBlogWithCategorys(fieldName string, fieldValue interface{}) (*models.Blog, error) {
	o := orm.NewOrm()
	if blog, err := GetBlog(fieldName, fieldValue); err != nil {
		return nil, err
	} else {
		_, err := o.LoadRelated(blog, "Categorys")
		return blog, err
	}
}

func GetCategory(fieldName string, fieldValue interface{}) (*models.Category, error) {
	o := orm.NewOrm()
	var category models.Category
	err := o.QueryTable("category").Filter(fieldName, fieldValue).One(&category)
	return &category, err
}

func CreateBlog(blog models.Blog) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&blog)
	return id, err
}

func CreateBlogWithCategorys(blog models.Blog, categorys []*models.Category) (int64, error) {
	o := orm.NewOrm()
	// todo Should create a transaction instead of this
	// insert blog
	id, err := o.Insert(&blog)
	// insert m2m
	m2m := o.QueryM2M(&blog, "Categorys")
	_, err = m2m.Add(categorys)
	if err != nil {
		return -1, err
	}
	return id, err
}

func CreateCategory(category models.Category) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&category)
	return id, err
}

func UpdateBlog(blog models.Blog) error {
	o := orm.NewOrm()
	_, err := o.Update(&blog)
	return err
}

func UpdateBlogWithCategory(blog models.Blog, categorys []*models.Category) error {
	o := orm.NewOrm()
	// todo Should create a transaction instead of this
	// delete and update m2m
	m2m := o.QueryM2M(&blog, "Categorys")
	// query exist m2m
	if _, err := o.LoadRelated(&blog, "Categorys"); err != nil {
		return err
	}
	// delete old m2m
	if len(blog.Categorys) != 0 {
		_, err := m2m.Remove(blog.Categorys)
		if err != nil {
			return err
		}
	}
	// add new m2m
	_, err := m2m.Add(categorys)
	if err != nil {
		return err
	}
	// update blog
	_, err = o.Update(&blog)
	return err
}

func UpdateCategory(category models.Category) error {
	o := orm.NewOrm()
	_, err := o.Update(&category)
	return err
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

func SearchBlog(search string, limit int) ([]models.Blog, error) {
	var blogs []models.Blog
	cond := orm.NewCondition()
	condition := cond.And("title__icontains", search).Or("content__icontains", search)
	qs := orm.NewOrm().QueryTable("blog")
	qs = qs.SetCond(condition)
	_, err := qs.Limit(limit).All(&blogs)
	return blogs, err
}