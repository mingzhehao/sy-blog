package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error404() {
	this.Data["content"] = "page not found"
	this.Data["Active"] = "error"
	this.Layout = "layout/default.html"
	this.TplName = "error/404.tpl"
}

func (this *ErrorController) Error501() {
	this.Data["Active"] = "error"
	this.Data["content"] = "server error"
	this.TplName = "error/501.tpl"
}

func (this *ErrorController) ErrorDb() {
	this.Data["Active"] = "error"
	this.Data["content"] = "database is now down"
	this.TplName = "error/dberror.tpl"
}
