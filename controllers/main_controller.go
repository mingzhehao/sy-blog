package controllers

import (
	"github.com/astaxie/beego"
	"github.com/mingzhehao/scloud/g"
	"github.com/mingzhehao/scloud/models"
	"github.com/mingzhehao/scloud/models/catalog"
	"strconv"
)

type MainController struct {
	BaseController
}

/**
 * 获取文章列表
 */
func (this *MainController) ArticleList() {
	currPage, _ := this.GetInt("p")
	if currPage == 0 {
		currPage = 1
	}
	pageSize, _ := strconv.Atoi(beego.AppConfig.String("pageSize"))
	total, _ := models.GetArticleCount()
	this.Data["Articles"], _, _ = models.GetArticles(currPage, pageSize)
	this.Data["HotArticles"], _, _ = models.GetHotArticles()
	this.SetPaginator(pageSize, total)
	beego.Notice(total)
	beego.Notice(this.Data["Articles"])
	this.Data["PageTitle"] = "首页"
	this.Data["Active"] = "list"
	this.Layout = "layout/default.html"
	this.TplName = "index.html"
}

/**
 * 获取目录列表
 */
func (this *MainController) CatalogList() {
	this.Data["Catalogs"] = catalog.All()
	this.Data["PageTitle"] = "首页"
	this.Data["Active"] = "catalog"
	this.Layout = "layout/default.html"
	this.TplName = "catalog/index.html"
}

func (this *MainController) Read() {
	ident := this.GetString(":ident")
	b := models.GetOneByIdent(ident)
	if b == nil {
		this.Ctx.WriteString("no such article")
		return
	}

	b.Views = b.Views + 1
	models.UpdateArticles(b, "")

	this.Data["Blog"] = b
	this.Data["Content"] = g.RenderMarkdown(models.ReadBlogContent(b).Content)
	this.Data["PageTitle"] = b.Title
	this.Data["Catalog"] = catalog.OneById(b.CatalogId)
	this.Data["Active"] = "list"
	this.Layout = "layout/default.html"
	this.TplName = "article/read.html"
}

func (this *MainController) ListByCatalog() {
	cata := this.Ctx.Input.Param(":ident")
	if cata == "" {
		this.Ctx.WriteString("catalog ident is blank")
		return
	}

	limit := this.GetIntWithDefault("limit", 10)

	c := catalog.OneByIdent(cata)
	if c == nil {
		this.Ctx.WriteString("catalog:" + cata + " not found")
		return
	}

	ids := models.GetIds(c.Id)
	pager := this.SetPaginator(limit, int64(len(ids)))
	articles := models.ByCatalog(c.Id, pager.Offset(), limit)

	this.Data["Catalog"] = c
	this.Data["Blogs"] = articles
	this.Data["PageTitle"] = c.Name
	this.Data["Active"] = "list"

	this.Layout = "layout/default.html"
	this.TplName = "article/by_catalog.html"
}
