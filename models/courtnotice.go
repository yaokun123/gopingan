package models

import (
	"strings"
	_"gopingan/models/dbresource/utnngrisk"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func GetCourtNoticesInfoByIds(ids []int)  {
	length := len(ids)
	where := strings.Repeat("?,",length)
	where = strings.TrimRight(where,",")
	where = "("+where+")"

	sql := "select * from court_notices where id in " + where + " order by startDate desc,id desc"
	O.Using("utn_ng_risk")

	var maps []orm.Params
	O.Raw(sql,ids).Values(&maps)

	for index,item := range maps {
		if item["startDate"] != nil{
			maps[index]["startDate"] = item["startDate"].(string)[0:10]
		}
		if item["update_date"] != nil{
			maps[index]["update_date"] = item["update_date"].(string)[0:10]
		}
		if item["create_time"] != nil{
			maps[index]["create_time"] = item["create_time"].(string)[0:10]
		}
	}

	fmt.Println(maps)
}


