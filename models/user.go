package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id       int64      `orm:"auto"`
	Name     string     `orm:"size(20)"`
	Bio      string
	Website  string
	Email    string
	Github   string
	Articles []*Article `orm:"reverse(many)"`
}

func InsertUser(a *User) error {
	o := orm.NewOrm()
	_, err := o.Insert(&a)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(a *User) error {
	o := orm.NewOrm()
	_, err := o.Delete(&a)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserById(id int64) error {
	o := orm.NewOrm()
	_, err := o.Delete(&User{Id: id})
	if err != nil {
		return err
	}
	return nil
}
