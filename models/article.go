package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	dir, _ := os.Getwd()
	confpath := filepath.Join(dir, "/../conf/app.conf")
	beego.LoadAppConfig("ini", confpath)
	usr := beego.AppConfig.String("dev::mysqluser")
	pass := beego.AppConfig.String("dev::mysqlpass")
	url := beego.AppConfig.String("dev::mysqlurls")
	db := beego.AppConfig.String("dev::mysqldb")
	ds := usr + ":" + pass + "@tcp(" + url + ")/" + db + "?charset=utf8"
	err := orm.RegisterDataBase("default", "mysql", ds, 1, 10)
	if err != nil {
		fmt.Println("RegisterDataBase Failed...")
	}
	orm.RegisterModel(new(Article))
	orm.RunSyncdb("default", false, true)
}

type Article struct {
	Id      int64     `orm:"auto" json:"id,string"`
	Title   string    `orm:"size(20)" json:"title"`
	Author  string    `orm:"size(20)" json:"author"`
	Created time.Time `orm:"auto_now_add" json:"-"`
	Updated time.Time `orm:"auto_now" json:"updated"`
	Content string    `orm:"size(65535)" json:"content"`
}

func (a *Article) TableIndex() [][]string {
	return [][]string{
		{"Id", "Title", "Author", "Updated"},
	}
}

func InsertArticle(a *Article) (int64, error) {
	o := orm.NewOrm()
	//TODO(luozhouyang): InsertOrUpdate does not work. Always insert a record and inc pk.
	_, err := o.InsertOrUpdate(a)
	if err != nil {
		return -1, err
	}
	err = o.Read(a, "Title", "Author", "Updated")
	if err != nil {
		return -1, err
	}
	return a.Id, nil
}

func DeleteArticleById(id int64) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Article{Id: id})
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticleByTitle(title string) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("article").Filter("title", title).Delete()
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticleByAuthor(author string) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("article").Filter("author", author).Delete()
	if err != nil {
		return err
	}
	return nil
}

func UpdateArticle(a *Article) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("article").Filter("id", a.Id).Update(orm.Params{
		"title":   a.Title,
		"author":  a.Author,
		"content": a.Content,
	})
	if err != nil {
		return err
	} else {
		return nil
	}
}

func QueryArticleById(id int64) (*Article, error) {
	o := orm.NewOrm()
	a := Article{Id: id}
	err := o.Read(&a, "Id")
	if err != nil {
		return nil, err
	} else {
		return &a, nil
	}
}

func QueryArticlesByTitle(title string) ([]*Article, error) {
	o := orm.NewOrm()
	var articles []*Article
	_, err := o.QueryTable("article").Filter("title", title).All(&articles)
	if err != nil {
		return nil, err
	} else {
		return articles, nil
	}
}

func QueryArticlesByAuthor(author string) ([]*Article, error) {
	o := orm.NewOrm()
	var articles []*Article
	_, err := o.QueryTable("article").Filter("author", author).All(&articles)
	if err != nil {
		return nil, err
	} else {
		return articles, nil
	}
}

type dateFormatErr struct {
	error
}

func (e dateFormatErr) Error() string {
	return "Format of date is incorrectly!"
}

func QueryArticlesByDate(date string) ([]*Article, error) {
	o := orm.NewOrm()
	days := strings.Split(date, "~")
	if len(days) != 2 {
		return nil, dateFormatErr{}
	}
	day0, day1 := days[0], days[1]
	var articles []*Article
	_, err := o.QueryTable("article").
		Filter("updated__gt", day0).Filter("updated__lt", day1).All(&articles)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func QueryArticlesByAuthorAndDateAndTitle(author, date, title string) ([]*Article, error) {
	o := orm.NewOrm()
	days := strings.Split(date, "~")
	if len(days) != 2 {
		return nil, dateFormatErr{}
	}
	day0, day1 := days[0], days[1]
	var articles []*Article
	_, err := o.QueryTable("article").Filter("author", author).
		Filter("updated__gt", day0).Filter("updated__lt", day1).Filter("title", title).All(&articles)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func QueryArticlesByAuthorAndDate(author, date string) ([]*Article, error) {
	o := orm.NewOrm()
	days := strings.Split(date, "~")
	if len(days) != 2 {
		return nil, dateFormatErr{}
	}
	day0, day1 := days[0], days[1]
	var articles []*Article
	_, err := o.QueryTable("article").Filter("author", author).
		Filter("updated__gt", day0).Filter("updated__lt", day1).All(&articles)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func QueryArticlesByAuthorAndTitle(author, title string) ([]*Article, error) {
	o := orm.NewOrm()
	var articles []*Article
	_, err := o.QueryTable("article").Filter("author", author).
		Filter("title", title).All(&articles)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func QueryArticlesByTitleAndDate(title, date string) ([]*Article, error) {
	o := orm.NewOrm()
	days := strings.Split(date, "~")
	if len(days) != 2 {
		return nil, dateFormatErr{}
	}
	day0, day1 := days[0], days[1]
	var articles []*Article
	_, err := o.QueryTable("article").
		Filter("updated__gt", day0).Filter("updated__lt", day1).
		Filter("title", title).All(&articles)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
