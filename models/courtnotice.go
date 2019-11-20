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

	sql := "select * from court_notices where id in "+where
	O.Using("utn_ng_risk")

	var maps []orm.Params
	O.Raw(sql,ids).Values(&maps)
	fmt.Println(maps)
}


