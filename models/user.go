package models

import (
	"time"
	"backend/db"
	"log"
	"fmt"
	"crypto/rand"
	"crypto/sha1"
)

type User struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) CreateSession() (session Session, err error) {
	statement := "INSERT INTO sessions (uuid, email, user_id, created_at) VALUES " +
		"($1, $2, $3, $4) RETURNING id, uuid, email, user_id, created_at"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), u.Email, u.Id, time.Now()).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

func (u *User) Session() (session Session, err error) {
	session = Session{}
	err = db.Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1", u.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
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

func (u *User) Create() (err error) {
	statement := "INSERT INTO USERS (uuid, name, email, password, created_at) VALUES " +
		"($1, $2, $3, $4, $5) RETURNING id, uuid, created_at"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now()).
		Scan(&u.Id, &u.Uuid, &u.CreatedAt)
	return
}

func (u *User) Delete() (err error) {
	statement := "DELETE FROM users WHERE id = $1"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Id)
	return
}

func (u *User) Update() (err error) {
	statement := "update users set name = $2, email = $3 where id = $1"
	stmt, err := db.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Id, u.Name, u.Email)
	return
}

func Users() (users []User, err error) {
	rows, err := db.Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	return
}

func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = db.Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = db.Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = $1", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
