package main

import (
	_ "backend/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	beego.BConfig.WebConfig.AutoRender = false
	usr := beego.AppConfig.String("dev::mysqluser")
	pass := beego.AppConfig.String("dev::mysqlpass")
	url := beego.AppConfig.String("dev::mysqlurls")
	db := beego.AppConfig.String("dev::mysqldb")
	ds := usr + ":" + pass + "@tcp(" + url + ")/" + db + "?charset=utf-8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", ds, 1, 10)
}

func main() {
	orm.Debug = true
	beego.Run()
}
