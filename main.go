package main

import (
	_ "backend/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	beego.BConfig.WebConfig.AutoRender = false
}

func main() {
	orm.Debug = true
	beego.Run()
}
