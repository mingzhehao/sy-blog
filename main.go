package main

import (
	"github.com/astaxie/beego"
	"github.com/mingzhehao/scloud/controllers"
	"github.com/mingzhehao/scloud/g"
	_ "github.com/mingzhehao/scloud/routers"
)

func main() {
	g.InitEnv()
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
