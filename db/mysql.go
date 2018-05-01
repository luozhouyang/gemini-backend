package db

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var ds string

func init() {
	config := &mysqlConfig{
		User:     "gemini",
		Password: "usergemini",
		DataBase: "gemini_db",
		Charset:  "utf8",
	}
	ds := config.DataSource()
	fmt.Printf("mysql datasource: %s", ds)
}

func NewDataBase() (*sql.DB, error) {
	db, err := sql.Open("mysql", ds)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

const CreateTable = "CREATE TABLE IF NOT EXISTS articles (" +
	"id INTEGER PRIMARY KEY AUTO_INCREMENT, " +
	"uuid VARCHAR(64) UNIQUE NOT NULL, " +
	"tags VARCHAR(255), " +
	"title VARCHAR(255) NOT NULL, " +
	"author VARCHAR(64) NOT NULL, " +
	"content TEXT NOT NULL, " +
	"updated_at DATE NULL DEFAULT NULL, " +
	"created_at DATE NULL DEFAULT NULL" +
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

type mysqlConfig struct {
	User     string
	Password string
	DataBase string
	Charset  string
}

func (c *mysqlConfig) DataSource() string {
	return fmt.Sprintf("%s:%s@localhost:3306/%s?charset=%s", c.User, c.Password, c.DataBase, c.Charset)
}
