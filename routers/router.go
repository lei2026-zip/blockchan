package routers

import (
	"demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/come_in", &controllers.ComeinController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
    beego.Router("/sign",&controllers.SignController{})
	beego.Router("/input",&controllers.InputController{})
}
