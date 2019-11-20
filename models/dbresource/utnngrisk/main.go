package utnngrisk

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init()  {
	dbhost := beego.AppConfig.String("utn_ng_risk_urls")
	dbport := beego.AppConfig.String("utn_ng_risk_port")
	dbuser := beego.AppConfig.String("utn_ng_risk_user")
	dbpassword := beego.AppConfig.String("utn_ng_risk_pass")
	dbname := beego.AppConfig.String("utn_ng_risk_db")

	if dbport == ""{
		dbport = "3306"
	}

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport +")" + "/" + dbname + "?charset=utf8"
	//orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("utn_ng_risk","mysql",dsn)

	sqldebug := beego.AppConfig.String("sqldebug")

	//是否开启数据库调试模式
	if sqldebug == "true"{
		orm.Debug=true
	}else{
		orm.Debug=false
	}

}
