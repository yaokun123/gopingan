package models

import (
	"github.com/astaxie/beego/orm"
	_"gopingan/models/dbresource/utndatanew"
	"time"
)

//开庭公告映射表
type CourtNoticesMap struct {
	Kid int
	Id int
	Company_name string
	StartDate time.Time `orm:"column(startDate)"`
}

var O orm.Ormer
func init()  {
	//开启调试模式
	orm.Debug=true


	orm.RegisterModel(new(CourtNoticesMap))
	O = orm.NewOrm()
}

func GetInfoByName(company_name string) []*CourtNoticesMap {
	var list []*CourtNoticesMap
	qs := O.QueryTable(new(CourtNoticesMap))
	qs.Filter("company_name",company_name).OrderBy("-startDate").All(&list)
	//qs.Filter("company_name",company_name).All(&list)
	return list
}