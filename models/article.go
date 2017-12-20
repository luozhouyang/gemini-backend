package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Article))
	orm.RunSyncdb(mysqldb, false, true)
}

type Article struct {
	Id       int64      `orm:"auto"`
	Title    string     `orm:"size(20)"`
	User     *User      `orm:"rel(fk)"`
	Time     time.Time
	Content  string
	Comments []*Comment `orm:"reverse(many)"`
}

func getOrmer() orm.Ormer {
	o := orm.NewOrm()
	o.Using(mysqldb)
	return o
}

func InsertArticle(a Article) error {
	o := getOrmer()
	_, err := o.Insert(a)
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticleById(id int64) error {
	o := getOrmer()
	_, err := o.Delete(&Article{Id: id})
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(a *Article) error {
	o := getOrmer()
	_, err := o.Delete(&a)
	if err != nil {
		return err
	}
	return nil
}
