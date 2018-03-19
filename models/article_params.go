package models

import (
	"encoding/json"
	"time"
)

type ArticleParam struct {
	Author string
	Title  string
	Date   string
}

func QueryArticlesJson(p *ArticleParam) (string, error) {
	date := concatDate(p.Date)
	if p.Author != "" && p.Title != "" && p.Date != "" {
		articles, err := QueryArticlesByAuthorAndDateAndTitle(p.Author, date, p.Title)
		return processArticles(articles, err)
	}
	if p.Author != "" && p.Title != "" {
		articles, err := QueryArticlesByAuthorAndTitle(p.Author, p.Title)
		return processArticles(articles, err)
	}
	if p.Author != "" && p.Date != "" {
		articles, err := QueryArticlesByAuthorAndDate(p.Author, date)
		return processArticles(articles, err)
	}
	if p.Title != "" && p.Date != "" {
		articles, err := QueryArticlesByTitleAndDate(p.Title, date)
		return processArticles(articles, err)
	}
	if p.Author != "" {
		articles, err := QueryArticlesByAuthor(p.Author)
		return processArticles(articles, err)
	}
	if p.Title != "" {
		articles, err := QueryArticlesByTitle(p.Title)
		return processArticles(articles, err)
	}
	if p.Date != "" {
		articles, err := QueryArticlesByDate(date)
		return processArticles(articles, err)
	}
	return "[]", nil
}

func processArticles(articles []*Article, err error) (string, error) {
	if err != nil {
		return "[]", err
	}
	jsons, err := json.Marshal(articles)
	if err != nil {
		return "[]", err
	}
	return string(jsons), nil
}

func concatDate(date string) string {
	if date == "" {
		return ""
	}
	today, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ""
	}
	tomorrow := today.Add(24 * time.Hour)
	date2 := today.String() + "~" + tomorrow.String()
	return date2
}

type ArticlePostParams struct {
	Title   string
	Author  string
	Updated string
	Content string
}

func PostArticle(p *ArticlePostParams) (int64, error) {
	updated := parseUpdatedTime(p.Updated)
	article := &Article{
		Title:   p.Title,
		Author:  p.Author,
		Updated: updated,
		Content: p.Content,
	}
	id, err := InsertOrUpdateArticle(article)
	return id, err
}

func parseUpdatedTime(updated string) time.Time {
	if updated == "" {
		return time.Now()
	}
	t, err := time.Parse("2006-01-02 00:00:00", updated)
	if err != nil {
		return time.Now()
	}
	return t
}
