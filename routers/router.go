package routers

import (
	"beeapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/post", &controllers.GreetingController{})
}
