package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
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

	//两个map
	uuid2PlaintiffsMap := make(map[string]string)
	uuid2defendantsMap := make(map[string]string)
	for _,val := range maps{
		uuid := val["lawsuit_uuid"].(string)
		uuid2PlaintiffsMap[uuid] = val["plaintiffs"].(string)
		uuid2defendantsMap[uuid] = val["defendants"].(string)
	}

	//再查原始数据
	var maps2 []orm.Params
	sql2 := "select * from company_lawsuit where uuid in " + where + " order by submittime desc,id desc"
	O.Raw(sql2,uuids).Values(&maps2)
	for index,item := range maps2 {
		//截取时间
		if item["submittime"] != nil{
			maps2[index]["submittime"] = item["submittime"].(string)[0:10]
		}
	}


	return maps2
}