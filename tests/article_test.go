package test

import (
	"testing"
	"backend/models"
	"runtime"
	"path/filepath"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath + "/../")
}

func TestInsertArticle(t *testing.T) {
	a := models.Article{
		Id:      1,
		Title:   "Hello World",
		Author:  "stupidme.me",
		Content: "Hello world from stupidme.me!",
	}
	err := models.InsertArticle(&a)
	if err != nil {
		//panic(err)
	}
	ar := models.QueryArticleById(1)
	if ar != nil {
		logs.GetLogger("TestArticle").Fatal("Insert article successfully!")
	} else {
		logs.GetLogger("TestArticle").Fatal("Insert article FAILED!")
	}
}
