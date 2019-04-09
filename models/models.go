package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Category struct {
	Id          int  `PK`
	Name        string  `orm:"size(20)"`
	Description string  `orm:"size(100);default('')"`
	Blogs        []*Blog  `orm:"reverse(many)"`
}

type User struct {
	Id              int  `PK`
	Username        string  `orm:"size(20)"`
	Password        string  `orm:"size(50)"`
	Blogs            []*Blog  `orm:"reverse(many)"`
}

type Blog struct {
	Id           int  `PK`
	Title        string  `orm:"size(100)"`
	PublicTime   time.Time  `orm:"auto_now_add;type(datetime)"`
	Author       *User  `orm:"rel(fk)"`
	Content      string  `orm:"size(5000)"`
	Categorys     []*Category  `orm:"rel(m2m);null"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "lgh:12345678@tcp(192.168.199.244:3306)/gosql?charset=utf8")
	//orm.RegisterDataBase("default", "mysql", "lgh:12345678@tcp(172.16.0.208:3306)/gosql?charset=utf8")
	orm.RegisterModel(new(User), new(Category), new(Blog))
	orm.RunSyncdb("default", false, true)

}