package db

import (
	"database/sql"
	"io/ioutil"
	"encoding/json"
	"log"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	f, err := ioutil.ReadFile("db/db.json")
	if err != nil {
		log.Fatal(err)
	}
	config := &config{}
	err = json.Unmarshal(f, config)
	if err != nil {
		log.Fatal(err)
	}
	ds := "host=" + config.Host + " port=" + config.Port + " user=" + config.Port + " password=" + config.Password +
		" dbname=" + config.DbName + " sslmode=" + config.SSLMode
	Db, err = sql.Open("postgres", ds)
	if err != nil {
		log.Fatal(err)
	}
}

type config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
}
