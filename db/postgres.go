package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"fmt"
)

var ds string

func init() {
	config := &config{
		Host:     "0.0.0.0",
		Port:     "8080",
		User:     "gemini",
		Password: "usergemini",
		DbName:   "geminidb",
		SSLMode:  "disable",
	}
	ds := config.DataSource()
	println(ds)
	err := createDatabase()
	if err != nil {
		panic(err)
	}
}

func NewDataBase() (*sql.DB, error) {
	db, err := sql.Open("postgres", ds)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

const CreateTable = "CREATE TABLE IF NOT EXISTS articles (" +
	"id SERIAL PRIMARY KEY, " +
	"uuid VARCHAR(64) UNIQUE NOT NULL, " +
	"tags VARCHAR(255), " +
	"title VARCHAR(255) NOT NULL, " +
	"author VARCHAR(255) NOT NULL, " +
	"content TEXT NOT NULL, " +
	"updated_at TIMESTAMP NOT NULL, " +
	"created_at TIMESTAMP NOT NULL" +
	")"

func createDatabase() error {
	db, err := NewDataBase()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(CreateTable)
	if err != nil {
		return err
	}
	return nil
}

type config struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  string
}

func (c *config) DataSource() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Port, c.DbName, c.SSLMode)
}
