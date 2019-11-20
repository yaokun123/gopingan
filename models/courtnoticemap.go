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
	orm.RegisterModel(new(CourtNoticesMap))
	O = orm.NewOrm()


	//开启调试模式
	orm.Debug=true
}

func GetCourtNoticesMapInfoByName(company_name string,offsetNum int) ([]CourtNoticesMap,int64) {
	var list []CourtNoticesMap
	qs := O.QueryTable(new(CourtNoticesMap))
	qs.Filter("company_name",company_name).OrderBy("-startDate","-id").Limit(10,offsetNum).All(&list)
	count,_ := qs.Filter("company_name",company_name).Count()

	return list,count
}