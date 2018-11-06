package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.DefaultString("db.port", "3306")
	dbUser := beego.AppConfig.String("db.user")
	dbPassword := beego.AppConfig.String("db.password")
	dbName := beego.AppConfig.String("db.name")
	//dbTimezone := beego.AppConfig.String("db.timezone")

	dbAddr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbName)
	//if dbTimezone != "" {
	//	dbAddr = dbAddr + "&loc=" + url.QueryEscape(dbTimezone)
	//}

	fmt.Println(dbAddr)
	orm.RegisterDataBase("default", "mysql", dbAddr)
	orm.RegisterModel(new(Article),new(InfoList),new(Admin))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	return  beego.AppConfig.String("db.name")+"."+
		beego.AppConfig.String("db.prefix") + name
}
