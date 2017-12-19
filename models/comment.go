package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Comment))
	orm.RunSyncdb(mysqldb, false, true)
}

type Comment struct {
	Id      int64  `orm:"auto"`
	Author  Author `orm:"reverse(one)"`
	Content string
}

func InsertComment(c *Comment) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(&c)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(c *Comment) (err error) {
	o := orm.NewOrm()
	_, err = o.Delete(&c)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCommentById(id int64) (err error) {
	o := orm.NewOrm()
	_, err = o.Delete(&Comment{Id: id})
	if err != nil {
		return err
	}
	return nil
}
