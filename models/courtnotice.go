package models

import (
	"strings"
	_"gopingan/models/dbresource/utnngrisk"
)

func GetCourtNoticesInfoByIds(ids []int)  {
	length := len(ids)
	where := strings.Repeat("?,",length)
	where = strings.TrimRight(where,",")
	where = "("+where+")"

	sql := "select * from court_notices where id in "+where
	O.Using("utn_ng_risk")
	list,_ := O.Raw(sql,ids).Exec()
	print(list)
}


