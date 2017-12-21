package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id       int64      `orm:"auto"`
	Title    string     `orm:"size(20)"`
	Author   *User      `orm:"rel(fk)"`
	Created  time.Time  `orm:"auto_now_add"`
	Updated  time.Time  `orm:"auto_now"`
	Content  string     `orm:"size(65535)"`
	Comments []*Comment `orm:"reverse(many)"`
}

func InsertArticle(a Article) error {
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
	_, err := o.Delete(&a)
	if err != nil {
		return err
	}
	return nil
}
