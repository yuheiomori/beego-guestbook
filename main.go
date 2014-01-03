package main

import (
	"beeapp/models"
	_ "beeapp/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"os"
	"strconv"
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

// テーブルがなければ作成する
func syncdb() {
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}
}

// データソース文字列を変換
func convert_datasource(ds string) (result string) {
	url, _ := url.Parse(ds)
	result = fmt.Sprintf("%s@tcp(%s:3306)%s", url.User.String(), url.Host, url.Path)
	beego.Info(result)
	return
}

func init() {

	var datasource string
	// for heroku with cleardb
	if os.Getenv("CLEARDB_DATABASE_URL") != "" {
		datasource = convert_datasource(os.Getenv("CLEARDB_DATABASE_URL"))
	} else {
		datasource = "user:pass@/database_name?charset=utf8"
	}
	orm.RegisterDataBase("default", "mysql", datasource, 30)
	orm.RegisterModel(new(models.Greeting))

	beego.AddFuncMap("dateformat", dateformat)
	beego.AddFuncMap("nl2br", nl2br)

}

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		beego.HttpPort = port
	}

	syncdb()
	beego.Run()
}
