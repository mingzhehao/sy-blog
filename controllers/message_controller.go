package controllers

import (
	"github.com/astaxie/beego"
	"github.com/mingzhehao/scloud/models"
	"html/template"
	"time"
)

type MessageController struct {
	BaseController
}

func (this *MessageController) MessageList() {
	var messages []*models.Message
	messages, _, _ = models.GetMessages(1, 10)
	this.Data["Messages"] = messages
	this.Data["Active"] = "message"
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.Layout = "layout/default.html"
	this.TplName = "message/index.html"
}

func (this *MessageController) AjaxAdd() {
	startTime := time.Now()
	if this.Ctx.Input.Method() != "POST" {
		this.Data["json"] = JsonFormat(101, "miss params", "", startTime)
		this.ServeJSON()
		return
	}
	uname := this.Input().Get("uname")
	content := this.Input().Get("content")
	if len(uname) == 0 || len(content) == 0 {
		this.Data["json"] = JsonFormat(102, "miss params", "", startTime)
		this.ServeJSON()
		return
	}
	//过滤参数

	//存储参数
	err := models.InsertMessage(0, 0, uname, content)
	if err != nil {
		beego.Notice(err)
		this.Data["json"] = JsonFormat(201, "insert err", "", startTime)
		this.ServeJSON()
	} else {
		this.Data["json"] = JsonFormat(1, "success", "", startTime)
		this.ServeJSON()
	}
	return
}
