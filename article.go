package main

import (
	"time"
	"backend/db"
)

type Article struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Title     string    `json:"title"`
	Tags      string    `json:"tags"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Article) Insert() error {
	database, err := db.NewDataBase()
	if err != nil {
		return err
	}
	defer database.Close()
	stmt, err := database.Prepare("INSERT INTO articles(uuid, title, tags, content, author, created_at, updated_at) " +
		"VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(a.Uuid, a.Title, a.Tags, a.Content, a.Author, a.CreatedAt, a.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (a *Article) Update() error {
	database, err := db.NewDataBase()
	if err != nil {
		return err
	}
	defer database.Close()
	stmt, err := database.Prepare("UPDATE articles SET title=?, tags=?, content=?, " +
		"author=?, updated_at=? WHERE uuid=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(a.Title, a.Tags, a.Content, a.Author, a.UpdatedAt, a.Uuid)
	if err != nil {
		return err
	}
	return nil
}

func (a *Article) Delete() error {
	database, err := db.NewDataBase()
	if err != nil {
		return err
	}
	defer database.Close()
	stmt, err := database.Prepare("DELETE * FROM articles WHERE uuid=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(a.Uuid)
	if err != nil {
		return err
	}
	return nil
}

func QueryByUuid(uuid string) (*Article, error) {
	database, err := db.NewDataBase()
	if err != nil {
		return nil, err
	}
	defer database.Close()
	rows, err := database.Query("SELECT id, uuid, title, tags, content, author, created_at, updated_at "+
		"FROM articles WHERE uuid=?", uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	a := Article{}
	rows.Scan(&a.Id, &a.Uuid, &a.Title, &a.Tags, &a.Content, &a.Author, &a.CreatedAt, &a.UpdatedAt)
	return &a, nil
}
