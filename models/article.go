package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"fmt"
)

func init() {
	beego.LoadAppConfig("ini", "/home/allen/Go/src/backend/conf/app.conf")
	usr := beego.AppConfig.String("dev::mysqluser")
	pass := beego.AppConfig.String("dev::mysqlpass")
	url := beego.AppConfig.String("dev::mysqlurls")
	db := beego.AppConfig.String("dev::mysqldb")
	ds := usr + ":" + pass + "@tcp(" + url + ")/" + db + "?charset=utf8"
	err := orm.RegisterDataBase("default", "mysql", ds, 1, 10)
	if err != nil {
		fmt.Println("RegisterDataBase Failed...")
	}
	orm.RegisterModel(new(Article))
	orm.RunSyncdb("default", false, true)
}

type Article struct {
	Id      int64     `orm:"auto"`
	Title   string    `orm:"size(20)"`
	Author  string    `orm:"size(20)"`
	Created time.Time `orm:"auto_now_add"`
	Updated time.Time `orm:"auto_now"`
	Content string    `orm:"size(65535)"`
}

func InsertArticle(a *Article) error {
	o := orm.NewOrm()
	_, err := o.Insert(a)
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticleById(id int64) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Article{Id: id})
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(a *Article) error {
	o := orm.NewOrm()
	_, err := o.Delete(a)
	if err != nil {
		return err
	}
	return nil
}

func UpdateArticle(a *Article) error {
	o := orm.NewOrm()
	_, err := o.Update(a)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func QueryArticleById(id int64) *Article {
	o := orm.NewOrm()
	a := Article{Id: id}
	err := o.Read(&a, "")
	if err != nil {
		return nil
	} else {
		return &a
	}
}

func QueryArticleByTitle(title string) *Article {
	o := orm.NewOrm()
	a := Article{Title: title}
	err := o.Read(&a, "")
	if err != nil {
		return nil
	} else {
		return &a
	}
}
