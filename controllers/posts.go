package controllers

import (
	"github.com/astaxie/beego"
	"backend/models"
)

type PostsController struct {
	beego.Controller
}

func (c *PostsController) Get() {
	p := &models.ArticleParam{
		Author: c.GetString("author"),
		Title:  c.GetString("title"),
		Date:   c.GetString("date"),
	}
	jsons, err := models.QueryArticlesJson(p)
	if err != nil {
		c.Data["json"] = "[]"
		c.ServeJSON(true, false)
		return
	}
	c.Data["json"] = jsons
	c.ServeJSON(true, false)
}

func (c *PostsController) Post() {

}
