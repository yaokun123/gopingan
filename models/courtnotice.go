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
		//maps[index]["startDate"] = item["startDate"]
		fmt.Println(index,item["startDate"])
	}
}


