package controllers

import (
	"github.com/astaxie/beego"
	"github.com/mingzhehao/scloud/models"
	"html/template"
	"strconv"
	"time"
)

type MessageController struct {
	BaseController
}

func (this *MessageController) MessageList() {
	var messages []*models.Message
	page, _ := this.GetInt("p")
	if page == 0 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(beego.AppConfig.String("pageSize"))
	total, _ := models.GetMessageCount()
	messages, _, _ = models.GetMessages(page, pageSize)
	beego.Notice(total)
	this.SetPaginator(pageSize, total)
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
