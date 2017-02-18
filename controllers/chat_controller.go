package controllers

import (
	"github.com/mingzhehao/scloud/chat"
)

type ChatController struct {
	BaseController
}

func (this *ChatController) WebSocket() {
	ControllerHub := chat.GlobalHub
	chat.ServeWs(ControllerHub, this.Ctx.ResponseWriter, this.Ctx.Request)
}

func (this *ChatController) Home() {
	this.Data["PageTitle"] = "聊天室"
	this.Data["Active"] = "chat"
	this.Layout = "layout/default.html"
	this.TplName = "chat/home.html"
}
