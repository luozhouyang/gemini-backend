package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Author))
	orm.RunSyncdb(mysqldb, false, true)
}

type Author struct {
	Id      int64  `orm:"auto"`
	Name    string `orm:"size(20)"`
	Bio     string
	Website string
	Email   string
	Github  string
}

func InsertAuthor(a *Author) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(&a)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAuthor(a *Author) (err error) {
	o := orm.NewOrm()
	_, err = o.Delete(&a)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAuthorById(id int64) (err error) {
	o := orm.NewOrm()
	_, err = o.Delete(&Author{Id: id})
	if err != nil {
		return err
	}
	return nil
}
