package routers

import (
	"beego-guestbook/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/post", &controllers.GreetingController{})
}
