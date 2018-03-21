package controllers

import (
	"github.com/astaxie/beego"
	"backend/models"
	"encoding/json"
	"backend/token"
)

type PostsController struct {
	beego.Controller
}

func (c *PostsController) Get() {
	p := &models.ArticleQueryParam{
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
	t := c.GetString("token")
	if t != token.GetToken() {
		c.Data["json"] = "[]"
		c.ServeJSON(true, false)
		return
	}
	p := &models.ArticlePostParams{
		Title:   c.GetString("title"),
		Author:  c.GetString("author"),
		Updated: c.GetString("updated"),
		Content: c.GetString("content"),
	}
	resp := models.ResponsePostArticle(p)
	jsons, err := json.Marshal(resp)
	if err != nil {
		c.Data["json"] = "[]"
		c.ServeJSON(true, false)
		return
	}
	c.Data["json"] = jsons
	c.ServeJSON()
}
