package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(Article))
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

func QueryArticleByTitle(title string) error {
	o := orm.NewOrm()
	a := Article{Title: title}
	err := o.Read(&a, "")
	if err != nil {
		return nil
	} else {
		return &a
	}
}
