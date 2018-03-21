package models

import (
	"encoding/json"
	"time"
)

type ArticleQueryParam struct {
	Author string
	Title  string
	Date   string
}

func QueryArticlesJson(p *ArticleQueryParam) (string, error) {
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

type ArticlePostResponse struct {
	ErrNumber int
	Message   string
	ArticleId int64
}

func ResponsePostArticle(p *ArticlePostParams) *ArticlePostResponse {
	updated := parseUpdatedTime(p.Updated)
	article := &Article{
		Title:   p.Title,
		Author:  p.Author,
		Updated: updated,
		Content: p.Content,
	}
	id, err := InsertOrUpdateArticle(article)
	if err != nil {
		return &ArticlePostResponse{
			ErrNumber: 1,
			Message:   "Error occurs: " + err.Error(),
			ArticleId: id,
		}
	}
	return &ArticlePostResponse{
		ErrNumber: 0,
		Message:   "Post article successfully.",
		ArticleId: id,
	}
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
