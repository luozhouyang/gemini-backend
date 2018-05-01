package main

import (
	"time"
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
