package controllers

import (
	"github.com/astaxie/beego"
	"github.com/mingzhehao/scloud/models"
	"github.com/mingzhehao/scloud/models/catalog"
	"strings"
)

type ArticleController struct {
	AdminController
}

func (this *ArticleController) Draft() {
	var articles []*models.Blog
	models.Blogs().Filter("Status", 0).All(&articles)
	this.Data["Blogs"] = articles
	this.Data["Active"] = "me"
	this.Layout = "layout/admin.html"
	this.TplName = "article/draft.html"
}

func (this *ArticleController) Add() {
	this.Data["Catalogs"] = catalog.All()
	this.Data["IsPost"] = true
	this.Data["Active"] = "me"
	this.Layout = "layout/admin.html"
	this.TplName = "article/add.html"
	this.JsStorage("deleteKey", "post/new")
}

func (this *ArticleController) DoAdd() {
	title := this.GetString("title")
	ident := this.GetString("ident")
	keywords := this.GetString("keywords")
	catalog_id := this.GetIntWithDefault("catalog_id", -1)
	aType := this.GetIntWithDefault("type", -1)
	status := this.GetIntWithDefault("status", -1)
	content := this.GetString("content")

	if catalog_id == -1 || aType == -1 || status == -1 {
		this.Ctx.WriteString("catalog || type || status is illegal")
		return
	}

	if title == "" || ident == "" {
		this.Ctx.WriteString("title or ident is blank")
		return
	}

	cp := catalog.OneById(int64(catalog_id))
	if cp == nil {
		this.Ctx.WriteString("catalog_id not exists")
		return
	}

	b := &models.Blog{Ident: ident, Title: title, Keywords: keywords, CatalogId: int64(catalog_id), Type: int8(aType), Status: int8(status)}
	blogId, err := models.SaveArticles(b, content)

	keyWords := strings.Split(keywords, ",")
	for _, tag := range keyWords {
		tag = strings.Trim(tag, " ")
		models.InsertTag(blogId, tag)
	}

	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.JsStorage("deleteKey", "post/new")
	this.Redirect("/catalog/"+cp.Ident, 302)

}

func (this *ArticleController) Edit() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.WriteString("get param id fail")
		return
	}

	b := models.GetOneById(int64(id))
	if b == nil {
		this.Ctx.WriteString("no such article")
		return
	}

	this.Data["Content"] = models.ReadBlogContent(b).Content
	this.Data["Blog"] = b
	this.Data["Catalogs"] = catalog.All()
	this.Data["Active"] = "me"
	this.Layout = "layout/admin.html"
	this.TplName = "article/edit.html"
}

func (this *ArticleController) DoEdit() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.WriteString("get param id fail")
		return
	}

	b := models.GetOneById(int64(id))
	if b == nil {
		this.Ctx.WriteString("no such article")
		return
	}

	title := this.GetString("title")
	ident := this.GetString("ident")
	keywords := this.GetString("keywords")
	catalog_id := this.GetIntWithDefault("catalog_id", -1)
	aType := this.GetIntWithDefault("type", -1)
	status := this.GetIntWithDefault("status", -1)
	content := this.GetString("content")

	if catalog_id == -1 || aType == -1 || status == -1 {
		this.Ctx.WriteString("catalog || type || status is illegal")
		return
	}

	if title == "" || ident == "" {
		this.Ctx.WriteString("title or ident is blank")
		return
	}
	oldKeyWordsString := b.Keywords

	cp := catalog.OneById(int64(catalog_id))
	if cp == nil {
		this.Ctx.WriteString("catalog_id not exists")
		return
	}

	b.Ident = ident
	b.Title = title
	b.Keywords = keywords
	b.CatalogId = int64(catalog_id)
	b.Type = int8(aType)
	b.Status = int8(status)

	err = models.UpdateArticles(b, content)

	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	/*******************************************/
	/**
	 * 标签替换处理
	 */
	var deleteKeyWords []string
	var addKeyWords []string
	oldKeyWords := strings.Split(oldKeyWordsString, ",")
	newKeyWords := strings.Split(keywords, ",")
	for _, oldWords := range oldKeyWords {
		//新数据包含老数据，说明无需处理
		//不包含，说明已删除
		if res := strings.Contains(keywords, oldWords); !res {
			deleteKeyWords = append(deleteKeyWords, oldWords)
		}
	}
	for _, newWords := range newKeyWords {
		if res := strings.Contains(oldKeyWordsString, newWords); !res {
			addKeyWords = append(addKeyWords, newWords)
		}
	}
	beego.Notice(addKeyWords)
	beego.Notice(deleteKeyWords)
	if len(addKeyWords) != 0 {
		for _, addTagString := range addKeyWords {
			addTagString = strings.Trim(addTagString, " ")
			models.InsertTag(b.Id, addTagString)
		}
	}
	if len(deleteKeyWords) != 0 {
		for _, deleteTagString := range deleteKeyWords {
			deleteTagString = strings.Trim(deleteTagString, " ")
			models.DeleteTag(b.Id, deleteTagString)
		}
	}

	/*******************************************/

	this.JsStorage("deleteKey", "post/edit")
	this.Redirect("/catalog/"+cp.Ident, 302)
}

func (this *ArticleController) Del() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.WriteString("get param id fail")
		return
	}

	b := models.GetOneById(int64(id))
	if b == nil {
		this.Ctx.WriteString("no such article")
		return
	}

	err = models.DelArticles(b)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.Ctx.WriteString("del success")
	return
}
