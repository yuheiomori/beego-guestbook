package controllers

import (
	"beeapp/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

// 投稿
type GreetingController struct {
	beego.Controller
}

func (this *GreetingController) Post() {
	o := orm.NewOrm()
	greet := models.Greeting{
		Name:     this.GetString("name"),
		Comment:  this.GetString("comment"),
		CreateAt: time.Now()}
	o.Insert(&greet)
	this.Ctx.Redirect(302, "/")
}
