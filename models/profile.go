package models

import (
	"time"
	"backend/db"
)

type Profile struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Bio       string    `json:"bio"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Profile) DeleteByUUID() (err error) {
	statement := "DELETE FROM profiles WHERE uuid = $1"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(&p.Uuid)
	return
}

func (p *Profile) User() (user User, err error) {
	statement := "SELECT FROM users WHERE id = $1"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(&p.UserId)
	return
}
