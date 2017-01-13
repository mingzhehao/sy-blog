package models

// package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Catalog struct {
	Id           int64
	Ident        string `orm:"unique"`
	Name         string
	Resume       string
	DisplayOrder int
	ImgUrl       string
}

type Blog struct {
	Id            int64
	Ident         string `orm:"unique"`
	Title         string
	Keywords      string       `orm:"null"`
	CatalogId     int64        `orm:"index"`
	Content       *BlogContent `orm:"-"`
	BlogContentId int64        `orm:"unique"`
	Type          int8         /*0:original, 1:translate, 2:reprint*/
	Up            int64
	Status        int8 /*0:draft, 1:release*/
	Views         int64
	Created       time.Time `orm:"auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"type(datetime)"`
}

type BlogContent struct {
	Id      int64
	Content string `orm:"type(text)"`
}

func (*Catalog) TableEngine() string {
	return engine()
}

func (*Blog) TableEngine() string {
	return engine()
}

func (*BlogContent) TableEngine() string {
	return engine()
}

func engine() string {
	return "INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci"
}

func init() {
	orm.RegisterModelWithPrefix("bb_", new(Catalog), new(Blog), new(BlogContent))
}

// func main() {
// 	orm.RegisterDataBase("default", "mysql", "root:@/beego_blog?charset=utf8&loc=Asia%2FShanghai", 30, 200)
// 	orm.RunCommand()
// }
