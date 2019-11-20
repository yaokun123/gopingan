package models

import (
	"strings"
	"fmt"
)

func GetCourtNoticesInfoByIds(ids []int)  {
	length := len(ids)
	where := strings.Repeat("?,",length)
	fmt.Println(where)
}


