package controllers

import (
	"github.com/mingzhehao/beego-blog/models/catalog"
)

type MeController struct {
	AdminController
}

func (this *MeController) Default() {
	this.Data["Active"] = "me"
	this.Data["IsList"] = true
	this.Data["Catalogs"] = catalog.All()
	this.Layout = "layout/admin.html"
	this.TplName = "me/default.html"
}
