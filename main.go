package main

import (
	"github.com/astaxie/beego"
	"github.com/mingzhehao/beego-blog/g"
	_ "github.com/mingzhehao/beego-blog/routers"
)

func main() {
	g.InitEnv()
	beego.Run()
}
