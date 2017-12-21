package test

import (
	"testing"
	"backend/models"
	"time"
	"runtime"
	"path/filepath"
	"github.com/astaxie/beego"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath + "/../")
}

func TestInsertArticle(t *testing.T) {
	comments := make([]*models.Comment, 1)
	a := models.Article{
		Id:      123456,
		Title:   "Hello World",
		Created: time.Now(),
		Author: &models.User{
			Id:      2314,
			Bio:     "",
			Name:    "Allen",
			Website: "",
			Github:  "https://github.com/luozhouyang",
		},
		Comments: comments,
		Content:  "Hello everyone! This is a test article.",
	}
	err := models.InsertArticle(a)
	if err != nil {
		t.Log("Error occurs when inserting article")
		return
	}
	t.Log("Successful to insert article")
}
