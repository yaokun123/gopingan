package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"fmt"
)

func GetCompanyLawsuitParsedInfoByUuids(uuids []string) []orm.Params {
	length := len(uuids)
	where := strings.Repeat("?,",length)
	where = strings.TrimRight(where,",")
	where = "(" + where + ")"

	sql := "select * from company_lawsuit_parsed_info where lawsuit_uuid in " + where
	O.Using("utn_ng_risk")

	var maps []orm.Params
	O.Raw(sql,uuids).Values(&maps)

	//uuid2PlaintiffsMap := make(map[string]string)
	//uuid2defendantsMap := make(map[string]string)

	for _,val := range maps{
		uuid := val["lawsuit_uuid"].(string)
		fmt.Println(uuid)
	}


	return maps
}