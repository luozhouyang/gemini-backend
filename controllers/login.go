package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type LoginController struct {
	beego.Controller
}

var log = logs.GetLogger("gemini-backend")

func (c *LoginController) Get() {
	log.Print("Login Get method called.")
}

func (c *LoginController) Post() {
	log.Print(c.Ctx.Request.PostForm)
}
