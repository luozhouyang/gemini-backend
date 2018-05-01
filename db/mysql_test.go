package db

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewDataBase(t *testing.T) {
	db, err := NewDataBase()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	assert.NotNil(t, db)
}

func TestMysqlConfig_DataSource(t *testing.T) {
	config := mysqlConfig{
		User:     "gemini",
		Password: "usergemini",
		DataBase: "gemini_db",
		Charset:  "utf8",
	}
	ds := config.DataSource()
	assert.Equal(t, "gemini:usergemini@localhost:3306/gemini_db?charset=utf8", ds)
}

