package models

import (
	"time"
	"backend/db"
)

type Comment struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Content   string    `json:"content"`
	UserId    int       `json:"user_id"`
	ArticleId int       `json:"article_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Comment) Create() (err error) {
	statement := "INSERT INTO comments (uuid, content, user_id, article_id, created_at " +
		"VALUES ($1, $2, $3, $4, $5) RETURNING id, uuid, created_at"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), c.Content, c.UserId, c.ArticleId, time.Now()).
		Scan(c.Id, c.Uuid, c.CreatedAt)
	return
}

func (c *Comment) Delete() (err error) {
	statement := "DELETE FROM comments WHERE id = $1"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(c.Id)
	return
}

func (u *User) CreateComment(c *Comment) (err error) {
	statement := "INSERT INTO comments (uuid, content, user_id, article_id, created_at " +
		"VALUES ($1, $2, $3, $4, $5) RETURNING id, uuid, created_at"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), c.Content, u.Id, c.ArticleId, time.Now()).
		Scan(c.Id, c.Uuid, c.CreatedAt)
	return
}

func (u *User) DeleteComment(c *Comment) (err error) {
	statement := "DELETE FROM comments WHERE id = $1 AND user_id = $2"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(c.Id, u.Id)
	return
}

func CommentsByArticleId(id int) (comments []Comment, err error) {
	rows, err := db.Db.Query("SELECT id, uuid, content, user_id, article_id, created_at FROM comments "+
		"WHERE article_id = $1", id)
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{}
		if err = rows.Scan(&comment.Id, &comment.Uuid, &comment.Content,
			&comment.UserId, &comment.ArticleId, &comment.CreatedAt); err != nil {
			return
		}
		comments = append(comments, comment)
	}
	return
}
