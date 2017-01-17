package main

import (
	"github.com/astaxie/beego"
	"github.com/mingzhehao/scloud/controllers"
	"github.com/mingzhehao/scloud/g"
	_ "github.com/mingzhehao/scloud/routers"
)

/**
 * 模板方法定义
 * 随机获取用户头像
 */
func GetUserImage(uid int64) (image string) {
	userImageArray := [4]string{"default1.jpg", "default2.jpg", "default3.jpg", "default4.jpg"}
	index := uid % 4
	return userImageArray[index]
}

func main() {
	g.InitEnv()
	beego.AddFuncMap("getUserImage", GetUserImage)
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
