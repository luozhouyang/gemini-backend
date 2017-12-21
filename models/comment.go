package models

import "github.com/astaxie/beego/orm"

type Comment struct {
	Id      int64    `orm:"auto"`
	Article *Article `orm:"rel(fk)"`
	User    *User    `orm:"rel(one)"`
	Content string   `orm:"size(1000)"`
}

func InsertComment(c *Comment) error {
	o := orm.NewOrm()
	_, err := o.Insert(c)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(c *Comment) error {
	o := orm.NewOrm()
	_, err := o.Delete(c)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCommentById(id int64) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Comment{Id: id})
	if err != nil {
		return err
	}
	return nil
}
