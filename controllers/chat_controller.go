package controllers

import (
	"fmt"
	"github.com/mingzhehao/scloud/chat"
)

type ChatController struct {
	BaseController
}

func (this *ChatController) WebSocket() {
	ControllerHub := chat.GlobalHub
	fmt.Println(ControllerHub)
	fmt.Println(this.Ctx.Request)
	chat.ServeWs(ControllerHub, this.Ctx.ResponseWriter, this.Ctx.Request)
}

func (this *ChatController) Home() {
	this.Data["Active"] = "chat"
	this.Layout = "layout/default.html"
	this.TplName = "chat/home.html"
}
