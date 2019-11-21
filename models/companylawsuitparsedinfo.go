package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"unicode/utf8"
	"gopingan/models/mymongo"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"fmt"
)

func GetCompanyLawsuitParsedInfoByUuids(uuids []string,digest string) []orm.Params {
	O := orm.NewOrm()
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

		//处理plaintiffs
		plaintiffs := uuid2PlaintiffsMap[item["uuid"].(string)]
		plaintiffsList := strings.Split(plaintiffs,",")
		var plaintiffsListResult []string
		for _,v := range plaintiffsList{
			arr := strings.Split(v,":")
			if utf8.RuneCountInString(arr[1]) > 4{//小于4个字的不取公司名
				digest := mymongo.GetMongoInfoByCompanyName(arr[1])
				if digest != ""{
					tmp := []string{arr[0],"<a target='_blank' class='col_link' href='/company-"+digest+".html' >"}
					plaintiffsListResult = append(plaintiffsListResult,strings.Join(tmp,":"))
				}else{
					plaintiffsListResult = append(plaintiffsListResult,strings.Join(arr,":"))
				}

			}
		}
		var plaintiffs_result string
		if len(plaintiffsListResult)>0{
			plaintiffs_result = strings.Join(plaintiffsListResult,",")
		}



		//处理defendants
		defendants := uuid2defendantsMap[item["uuid"].(string)]
		defendantsList := strings.Split(defendants,",")
		var defendantsListResult []string
		for _,v := range defendantsList{
			arr := strings.Split(v,":")
			if utf8.RuneCountInString(arr[1]) > 4{//小于4个字的不取公司名
				digest := mymongo.GetMongoInfoByCompanyName(arr[1])
				if digest != ""{
					tmp := []string{arr[0],"<a target='_blank' class='col_link' href='/company-"+digest+".html' >"}
					defendantsListResult = append(defendantsListResult,strings.Join(tmp,":"))
				}else{
					defendantsListResult = append(defendantsListResult,strings.Join(arr,":"))
				}

			}
		}
		var defendants_result string
		if len(defendantsListResult) >0{
			defendants_result = strings.Join(defendantsListResult,",")
		}


		if plaintiffs_result != "" &&defendants_result != "" {
			maps2[index]["case_position"] = plaintiffs_result + "," + defendants_result
		}else if plaintiffs_result != ""{
			maps2[index]["case_position"] = plaintiffs_result
		}else if defendants_result != ""{
			maps2[index]["case_position"] = defendants_result
		}


		//签名
		h := md5.New()
		h.Write([]byte(item["id"].(string)))
		cipherStr := h.Sum(nil)
		idSign := hex.EncodeToString(cipherStr)

		pre := idSign[0:6]//前6位
		next := idSign[26:]//后6位
		intId,_:=strconv.Atoi(item["id"].(string))
		tmp_id := pre + strconv.Itoa(intId+12345) + next

		has := md5.Sum([]byte(tmp_id + "_page_courts_new"))
		md5str := fmt.Sprintf("%x",has)
		maps2[index]["id_sign"] = md5str

		//处理一下title
		maps2[index]["title"] = "<a class='col_link' target='_blank' href='/page-courts-detail?id="+tmp_id+"&id_sign="+md5str+"&digest="+digest+"' >" + item["title"].(string) + "</a>"
	}


	return maps2
}