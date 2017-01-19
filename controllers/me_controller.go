package controllers

import (
	"github.com/astaxie/beego"
	"github.com/mingzhehao/scloud/models"
	"strconv"
)

type MeController struct {
	AdminController
}

func (this *MeController) Default() {
	currPage, _ := this.GetInt("p")
	if currPage == 0 {
		currPage = 1
	}
	pageSize, _ := strconv.Atoi(beego.AppConfig.String("pageSize"))
	this.Data["Active"] = "me"
	this.Data["IsList"] = true
	this.Data["Catalogs"], _, _ = models.GetCatalogs(currPage, pageSize)
	this.Layout = "layout/admin.html"
	this.TplName = "me/default.html"
}
