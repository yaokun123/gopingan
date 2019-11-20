package utndatanew

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
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

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport +")" + "/" + dbname + "?charset=utf8"
	//orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",dsn)

	//设置数据库链接池
	maxOpenConns,_ := beego.AppConfig.Int("maxopenconns")
	maxIdleConns,_ := beego.AppConfig.Int("maxidleconns")
	orm.SetMaxOpenConns("default",maxOpenConns)
	orm.SetMaxIdleConns("default",maxIdleConns)


	//是否开启数据库调试模式
	sqldebug := beego.AppConfig.String("sqldebug")
	if sqldebug == "true"{
		orm.Debug=true
	}else{
		orm.Debug=false
	}

}
