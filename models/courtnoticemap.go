package models

import (
	"github.com/astaxie/beego/orm"
	_"gopingan/models/dbresource/utndatanew"
	"time"
)

//开庭公告映射表
type CourtNoticesMap struct {
	Kid int `orm:"description(注释这么写)"`
	Id int
	Company_name string
	StartDate time.Time `orm:"column(startDate);description(注释这么写)"`
}


func init()  {
	orm.RegisterModel(new(CourtNoticesMap))
}

func GetCourtNoticesMapInfoByName(company_name string,offsetNum int) ([]CourtNoticesMap,int64) {
	O := orm.NewOrm()
	var list []CourtNoticesMap
	O.Using("default")
	qs := O.QueryTable(new(CourtNoticesMap))
	qs.Filter("company_name",company_name).OrderBy("-startDate","-id").Limit(5,offsetNum).All(&list)
	count,_ := qs.Filter("company_name",company_name).Count()

	return list,count
}