package main

import (
	"beeapp/models"
	_ "beeapp/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
	"time"
)

// time.Timeオブジェクトを見やすい表示にする関数
func dateformat(in time.Time) (out string) {
	out = in.Format("2006-01-02 15:04:05")
	return
}

// 改行文字をbrタグに置き換える関数
func nl2br(in string) (out string) {
	out = strings.Replace(in, "\n", "<br>", -1)
	return
}

func init() {
	orm.RegisterDataBase("default", "mysql", os.Getenv("BEEAPP_DB"), 30)
	orm.RegisterModel(new(models.Greeting))

	beego.AddFuncMap("dateformat", dateformat)
	beego.AddFuncMap("nl2br", nl2br)

}

// テーブルがなければ作成する
func syncdb() {
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}
}

func main() {
	syncdb()
	beego.Run()
}
