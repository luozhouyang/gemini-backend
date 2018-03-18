package controllers

import (
	"github.com/astaxie/beego"
	"backend/models"
	"encoding/json"
	"time"
	"log"
)

type PostsController struct {
	beego.Controller
}

func (c *PostsController) Get() {
	author := c.GetString("author")
	date := c.GetString("date")
	today, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatal(err)
		return
	}
	tomorrow := today.Add(24 * time.Hour)
	date = today.String() + "~" + tomorrow.String()
	log.Println(date)
	title := c.GetString("title")
	if author != "" && date != "" && title != "" {
		articles, err := models.QueryArticlesByAuthorAndDateAndTitle(author, date, title)
		serverArticles(c, articles, err)
		return
	}
	if author != "" && date != "" {
		articles, err := models.QueryArticlesByAuthorAndDate(author, date)
		serverArticles(c, articles, err)
		return
	}
	if author != "" && title != "" {
		articles, err := models.QueryArticlesByAuthorAndTitle(author, title)
		serverArticles(c, articles, err)
		return
	}
	if date != "" && title != "" {
		articles, err := models.QueryArticlesByTitleAndDate(title, date)
		serverArticles(c, articles, err)
		return
	}
	if author != "" {
		articles, err := models.QueryArticlesByAuthor(author)
		serverArticles(c, articles, err)
		return
	}
	if title != "" {
		articles, err := models.QueryArticlesByTitle(title)
		serverArticles(c, articles, err)
		return
	}
	if date != "" {
		articles, err := models.QueryArticlesByDate(date)
		serverArticles(c, articles, err)
		return
	}
	c.Data["json"] = "[]"
	c.ServeJSON(true, false)
}

func serverArticles(c *PostsController, articles []*models.Article, err error) {
	if err != nil {
		c.Data["json"] = "[]"
		c.ServeJSON(true, false)
		return
	}
	jsons, err := json.Marshal(articles)
	if err != nil {
		c.Data["json"] = "[]"
		c.ServeJSON(true, false)
		return
	}
	c.Data["json"] = string(jsons)
	c.ServeJSON(true, false)
	return
}

func (c *PostsController) Post() {

}
