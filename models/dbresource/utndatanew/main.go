package utndatanew

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func init()  {
	dbhost := beego.AppConfig.String("utn_data_new_urls")
	dbport := beego.AppConfig.String("utn_data_new_port")
	dbuser := beego.AppConfig.String("utn_data_new_user")
	dbpassword := beego.AppConfig.String("utn_data_new_pass")
	dbname := beego.AppConfig.String("utn_data_new_db")

	if dbport == ""{
		dbport = "3306"
	}

	dsn := dbuser + ":" + dbpassword + "@" + dbhost + ":" + dbport + "/" + dbname + "?charset=utf8"
	fmt.Println(dsn)
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",dsn)
}
