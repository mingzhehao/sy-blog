package controllers

import (
	"strings"
)

/**
 * 定义tools静态页面
 */
const ToolsString = "git,svn,php,go,python,jquery,java,python,html,nodejs,ruby,javascript,css,svn,mysql,regex,linux,html,lua,html-dom,express"

type ToolsController struct {
	BaseController
}

func (this *ToolsController) ToolsList() {
	this.Data["Active"] = "tools"
	this.Layout = "layout/main.html"
	this.TplName = "tools/index.html"
}

func (this *ToolsController) Read() {
	ident := this.GetString(":ident")
	if len(ident) == 0 {
		this.Abort("404")
		return
	}
	if exists := strings.Contains(ToolsString, ident); !exists {
		this.Abort("404")
		return
	}
	this.Data["Active"] = "tools"
	this.Layout = "layout/main.html"
	this.TplName = "tools/" + ident + ".html"
}
