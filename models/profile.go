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

func (u *User) CreateProfile(bio string) (profile Profile, err error) {
	statement := "INSERT INTO profiles (uuid, bio, user_id, created_at) " +
		"VALUES ($1, $2, $3, $4) RETURNING id, uuid, bio, user_id, created_at"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), bio, u.Id, time.Now()).
		Scan(&profile.Id, &profile.Uuid, &profile.Bio, &profile.UserId, &profile.CreatedAt)
	return
}

func (u *User) Profile() (profile Profile, err error) {
	profile = Profile{}
	err = db.Db.QueryRow("SELECT id, uuid, bio, user_id, created_at FROM profiles WHERE user_id = $1", u.Id).
		Scan(&profile.Id, &profile.Uuid, &profile.Bio, &profile.UserId, &profile.CreatedAt)
	return
}
