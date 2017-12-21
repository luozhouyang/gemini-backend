package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

func init() {
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlurl := beego.AppConfig.String("mysqlurls")
	mysqldbname := beego.AppConfig.String("mysqldb")
	datasource := mysqluser + ":" + mysqlpass + "@tcp(" + mysqlurl + ")/" + mysqldbname + "?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	datasource = "gemini:usergemini@tcp(localhost:3306)/gemini_db?charset=utf8"
	orm.RegisterDataBase("default", "mysql", datasource)
	orm.RegisterModel(new(Article), new(User), new(Comment))
	orm.RunSyncdb("default", false, true)
}
