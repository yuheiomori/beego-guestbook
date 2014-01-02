package models

import (
	"time"
)

type Greeting struct {
	Id       int    `orm:"auto"`
	Name     string `orm:"size(100)"`
	Comment  string `orm:"size(200)"`
	CreateAt time.Time
}
