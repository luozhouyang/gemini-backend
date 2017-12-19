package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Comment))
	orm.RunSyncdb(mysqldb, false, true)
}

type Comment struct {
	Id      int64 `orm:"auto"`
	User    User  `orm:"reverse(one)"`
	Content string
}

func InsertComment(c *Comment) error {
	o := orm.NewOrm()
	_, err := o.Insert(&c)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(c *Comment) error {
	o := orm.NewOrm()
	_, err := o.Delete(&c)
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
