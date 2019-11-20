package models

import (
	"strings"
)

func GetCourtNoticesInfoByIds(ids []int)  {
	length := len(ids)
	where := strings.Repeat("?,",length)
	where = strings.TrimRight(where,",")
	where = "("+where+")"

	sql := "select * form court_notices where id in "+where
	list,_ := O.Raw(sql,ids).Exec()
	print(list)
}


