package routers

import (
	"star/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.UserController{})
	beego.Router("/category",&controllers.UserController{},`get:Category`)
	beego.Router("/api/user/profile",&controllers.UserController{},`get:API_Profile`)
	beego.Router("/login",&controllers.LoginController{},`post:Login`)
	beego.Router("/logout",&controllers.LoginController{},`get:Logout`)
	beego.Router("/register",&controllers.UserController{},`post:Register`)
	beego.Router("/setting",&controllers.UserController{},`post:Setting;get:Pageinfo`)
	
}
