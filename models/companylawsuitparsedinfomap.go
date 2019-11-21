package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

//法律诉讼映射表
type  CompanyLawsuitParsedInfoMap struct {
	Kid int
	Id int
	Company_name string
	Submittime time.Time
	Uuid string
	Casereason string
}

func init()  {
	orm.RegisterModel(new(CompanyLawsuitParsedInfoMap))
}

func GetCompanyLawsuitParsedInfoMapInfoByName(company_name string,offsetNum int) ([]CompanyLawsuitParsedInfoMap,int64) {
	O := orm.NewOrm()
	var list []CompanyLawsuitParsedInfoMap
	O.Using("default")
	qs := O.QueryTable(new(CompanyLawsuitParsedInfoMap))
	qs.Filter("company_name",company_name).OrderBy("-submittime","-id").Limit(5,offsetNum).All(&list)
	count,_ := qs.Filter("company_name",company_name).Count()
	return list,count
}