package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

func init() {
	mysqluser := beego.AppConfig.DefaultString("mysqluser", "gemini")
	mysqlpass := beego.AppConfig.DefaultString("mysqlpass", "usergemini")
	mysqlurl := beego.AppConfig.DefaultString("mysqlurls", "localhost:3306")
	mysqldbname := beego.AppConfig.DefaultString("mysqldb", "gemini_db")
	datasource := mysqluser + ":" + mysqlpass + "@tcp(" + mysqlurl + ")/" + mysqldbname + "?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//datasource = "gemini:usergemini@tcp(localhost:3306)/gemini_db?charset=utf8"
	orm.RegisterDataBase("default", "mysql", datasource)
	orm.RegisterModel(new(Article), new(User), new(Comment))
	orm.RunSyncdb("default", false, true)
}
