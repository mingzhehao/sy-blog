package routers

import (
	"github.com/astaxie/beego"
	"github.com/mingzhehao/scloud/controllers"
)

func init() {

	beego.AutoRouter(&controllers.ApiController{})

	beego.Router("/", &controllers.MainController{}, "get:ArticleList")
	beego.Router("/catalog", &controllers.MainController{}, "get:CatalogList")
	beego.Router("/article/:ident", &controllers.MainController{}, "get:Read")
	beego.Router("/catalog/:ident", &controllers.MainController{}, "get:ListByCatalog")

	beego.Router("/login", &controllers.LoginController{}, "get:Login;post:DoLogin")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	beego.Router("/me", &controllers.MeController{}, "get:Default")
	beego.Router("/me/catalog/add", &controllers.CatalogController{}, "get:Add;post:DoAdd")
	beego.Router("/me/catalog/edit", &controllers.CatalogController{}, "get:Edit;post:DoEdit")
	beego.Router("/me/catalog/del", &controllers.CatalogController{}, "get:Del")
	beego.Router("/me/article/add", &controllers.ArticleController{}, "get:Add;post:DoAdd")
	beego.Router("/me/article/edit", &controllers.ArticleController{}, "get:Edit;post:DoEdit")
	beego.Router("/me/article/del", &controllers.ArticleController{}, "get:Del")
	beego.Router("/me/article/draft", &controllers.ArticleController{}, "get:Draft")

	/*留言板*/
	beego.Router("/message", &controllers.MessageController{}, "get:MessageList")
	beego.Router("/message/add", &controllers.MessageController{}, "post:AjaxAdd")

	/*工具箱*/
	beego.Router("/tools", &controllers.ToolsController{}, "get:ToolsList")
	beego.Router("/tools/:ident", &controllers.ToolsController{}, "get:Read")

	/*聊天室*/
	beego.Router("/chat", &controllers.ChatController{}, "get:Home")
	beego.Router("/ws", &controllers.ChatController{}, "get:WebSocket")
}
