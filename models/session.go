package models

import (
	"time"
	"backend/db"
)

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func (s *Session) Valid() (valid bool, err error) {
	err = db.Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", s.Uuid).
		Scan(&s.Id, &s.Uuid, &s.Email, &s.UserId, &s.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if s.Id != 0 {
		valid = true
	}
	return
}

func (s *Session) DeleteByUUID(uuid string) (err error) {
	statement := "DELETE FROM sessions WHERE uuid = $1"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(s.Uuid)
	return
}

func (s *Session) User(user User, err error) {
	user = User{}
	err = db.Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", s.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}
