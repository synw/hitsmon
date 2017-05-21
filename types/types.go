package types

import (
	"time"
)

type Conf struct {
	DbType    string
	Addr      string
	User      string
	Pwd       string
	Db        string
	Table     string
	Frequency int
	Separator string
	Dev       bool
	Verb      int
}

type Hit struct {
	Id        uint `gorm:"primary_key"`
	Domain    string
	Path      string
	Method    string
	Ip        string
	UserAgent string `gorm:"column:user_agent"`
	User      string
	Referer   string
	Date      time.Time
}
