package models

import (
	"time"
	"backend/db"
)

type Article struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Tags      string    `json:"tags"`
	Content   string    `json:"content"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) CreateArticle(a *Article) (err error) {
	statement := "INSERT INTO articles (uuid, tags, content, user_id, created_at, updated_at) " +
		"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, uuid, created_at"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	t := time.Now()
	err = stmt.QueryRow(createUUID(), a.Tags, a.Content, u.Id, t, t).
		Scan(&a.Id, &a.Uuid, &a.CreatedAt)
	return
}

func (u *User) UpdateArticle(a *Article) (err error) {
	statement := "UPDATE articles SET tags = $2, content = $3, updated_at = $4 WHERE id = $1 AND user_id = $5"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(a.Id, a.Tags, a.Content, time.Now(), u.Id)
	return
}

func (u *User) DeleteArticle(a *Article) (err error) {
	statement := "DELETE FROM articles WHERE id = $1 AND user_id = $2"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return nil
	}
	defer stmt.Close()
	_, err = stmt.Exec(a.Id, u.Id)
	return
}

//TODO(luozhouyang) operate articles by session, which goal needs an efficient memory session manager

func (u *User) ArticleById(id int) (article Article, err error) {
	statement := "SELECT id, uuid, tags, content, user_id, created_at, updated_at FROM articles WHERE id = $1"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	article = Article{}
	err = stmt.QueryRow(id).Scan(&article.Id, &article.Uuid, &article.Tags, &article.Content,
		&article.UserId, &article.CreatedAt, &article.UpdatedAt)
	return
}

func (u *User) ArticleByUUID(uuid string) (article Article, err error) {
	statement := "SELECT id, uuid, tags, content, user_id, created_at, updated_at FROM articles WHERE uuid = $1"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	article = Article{}
	err = stmt.QueryRow(uuid).Scan(&article.Id, &article.Uuid, &article.Tags, &article.Content,
		&article.UserId, &article.CreatedAt, &article.UpdatedAt)
	return
}

func (u *User) ArticlesByTag(tags string) (articles []Article, err error) {
	statement := "SELECT id, uuid, tags, content, user_id, created_at, updated_at FROM articles " +
		"WHERE tags LIKE CONCAT('%', $1, '%')"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(tags)
	if err != nil {
		return
	}
	for rows.Next() {
		article := Article{}
		if err = rows.Scan(&article.Id, &article.Uuid, &article.Tags, &article.Content,
			&article.UserId, &article.CreatedAt, &article.UpdatedAt); err != nil {
			return
		}
		articles = append(articles, article)
	}
	return
}

func (u *User) ArticlesOfUser() (articles []Article, err error) {
	statement := "SELECT id, uuid, tags, content, user_id, created_at, updated_at FROM articles " +
		"WHERE user_id = $1"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(u.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		article := Article{}
		if err = rows.Scan(&article.Id, &article.Uuid, &article.Tags, &article.Content,
			&article.UserId, &article.CreatedAt, &article.UpdatedAt); err != nil {
			return
		}
		articles = append(articles, article)
	}
	return
}

//TODO(luozhouyang) inflate these methods
func (u *User) ArticlesByDate(date string) (articles []Article, err error) {
	return
}

func (u *User) ArticlesByMonth(month string) (articles []Article, err error) {
	return
}

func (u *User) ArticlesByYear(year string) (articles []Article, err error) {
	return
}

func (u *User) Articles() (articles []Article, err error) {
	return
}
