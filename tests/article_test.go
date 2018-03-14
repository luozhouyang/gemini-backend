package test

import (
	"testing"
	"backend/models"
	"runtime"
	"path/filepath"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath + "/../")
}

func TestInsertArticle(t *testing.T) {
	a := models.Article{
		Id:      5,
		Title:   "Hello Worldddd",
		Author:  "stupidme.me",
		Content: "Hello world from stupidme.me! Hello",
	}
	err := models.InsertArticle(&a)
	if err != nil {
		panic(err)
	}
	logs.GetLogger("test_").Println("Insert article successfully.")
}

func TestUpdateArticle(t *testing.T) {
	a := models.Article{
		Id:      5,
		Title:   "Hello",
		Author:  "stupidme.me",
		Content: "Hello world from stupidme.me! Hello",
	}
	err := models.UpdateArticle(&a)
	if err != nil {
		panic(err)
	}
}

func TestDeleteArticleById(t *testing.T) {
	err := models.DeleteArticleById(1)
	if err != nil {
		panic(err)
	}
}

func TestDeleteArticleByTitle(t *testing.T) {
	err := models.DeleteArticleByTitle("Hello World")
	if err != nil {
		panic(err)
	}
}

func TestDeleteArticleByAuthor(t *testing.T) {
	err := models.DeleteArticleByAuthor("stupidme.me")
	if err != nil {
		panic(err)
	}
}

func TestQueryArticleById(t *testing.T) {
	article, err := models.QueryArticleById(3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", *article)
}

func TestQueryArticlesByAuthor(t *testing.T) {
	articles, err := models.QueryArticlesByAuthor("stupidme.me")
	if err != nil {
		panic(err)
	}
	for i, v := range articles {
		fmt.Printf("%d-%v\n", i, *v)
	}
}

func TestQueryArticlesByTitle(t *testing.T) {
	articles, err := models.QueryArticlesByTitle("Hello World")
	if err != nil {
		panic(err)
	}
	for i, v := range articles {
		fmt.Printf("%d-%v\n", i, *v)
	}
}
