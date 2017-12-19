package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

func init() {
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlurl := beego.AppConfig.String("mysqlurl")
	mysqldb := beego.AppConfig.String("mysqldb")
	datasource := mysqluser + ":" + mysqlpass + "@tcp(" + mysqlurl + ")/" + mysqldb + "?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", datasource)
}

var mysqldb = beego.AppConfig.String("mysqldb")
