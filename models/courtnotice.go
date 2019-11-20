package models

import (
	"strings"
	"fmt"
)

func GetCourtNoticesInfoByIds(ids []int)  {
	length := len(ids)
	where := strings.Repeat("?,",length)
	where = where[0:-1]
	fmt.Println(where)
}


