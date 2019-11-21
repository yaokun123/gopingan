package models

import (
	"strings"
	_"gopingan/models/dbresource/utnngrisk"
	"github.com/astaxie/beego/orm"
	"unicode/utf8"
	"gopingan/models/mymongo"
)

func GetCourtNoticesInfoByIds(ids []int) []orm.Params {
	length := len(ids)
	where := strings.Repeat("?,",length)
	where = strings.TrimRight(where,",")
	where = "("+where+")"

	sql := "select * from court_notices where id in " + where + " order by startDate desc,id desc"
	O.Using("utn_ng_risk")

	var maps []orm.Params
	O.Raw(sql,ids).Values(&maps)

	//公司名检测plaintiff/defendant
	var company_name_list []string
	for index,item := range maps {

		//截取时间
		if item["startDate"] != nil{
			maps[index]["startDate"] = item["startDate"].(string)[0:10]
		}
		if item["update_date"] != nil{
			maps[index]["update_date"] = item["update_date"].(string)[0:10]
		}
		if item["create_time"] != nil{
			maps[index]["create_time"] = item["create_time"].(string)[0:10]
		}
		//处理公诉人/原告/上诉人/申请人
		if item["plaintiff"] != nil{
			plaintiffs := strings.Split(item["plaintiff"].(string),",")
			for index2,company_name := range plaintiffs{
				if utf8.RuneCountInString(company_name) > 4{//小于4个字的不取公司名
					company_name_list = append(company_name_list,company_name)
				}else if utf8.RuneCountInString(company_name) == 3{
					tmp := []rune(company_name)
					first_char := string(tmp[0:1])
					end_char := string(tmp[2:])
					result_name := first_char + "*" + end_char
					plaintiffs[index2] = result_name
				}else if utf8.RuneCountInString(company_name) == 2{
					tmp := []rune(company_name)
					first_char := string(tmp[0:1])
					result_name := first_char + "*"
					plaintiffs[index2] = result_name
				}
			}
			maps[index]["plaintiff"] = plaintiffs
		}

		//处理被告人/被告/被上诉人/被申请人
		if item["defendant"] != nil{
			defendants := strings.Split(item["defendant"].(string),",")
			for index2,company_name := range defendants{
				if utf8.RuneCountInString(company_name) > 4{//小于4个字的不取公司名
					company_name_list = append(company_name_list,company_name)
				}else if utf8.RuneCountInString(company_name) == 3{
					tmp := []rune(company_name)
					first_char := string(tmp[0:1])
					end_char := string(tmp[2:])
					result_name := first_char + "*" + end_char
					defendants[index2] = result_name
				}else if utf8.RuneCountInString(company_name) == 2{
					tmp := []rune(company_name)
					first_char := string(tmp[0:1])
					result_name := first_char + "*"
					defendants[index2] = result_name
				}
			}
			maps[index]["defendant"] = defendants
		}

	}

	//获取company_name_2_digest的map
	var company_name_2_digest_map map[string]string
	for _,v := range company_name_list{
		company_name_2_digest_map[v] = mymongo.GetMongoInfoByCompanyName(v)
	}

	//最后再处理一次数据，将digest加上
	/*for index,item := range maps{
		var plaintiffDigests []string
		if item["plaintiff"] != nil{
			for _,v := range item["plaintiff"].([]string){
				plaintiffDigests = append(plaintiffDigests,company_name_2_digest_map[v])
			}
		}
		maps[index]["plaintiff_digest"] = plaintiffDigests

		var defendantDigests []string
		if item["defendant"] != nil{
			for _,v := range item["defendant"].([]string){
				defendantDigests = append(defendantDigests,company_name_2_digest_map[v])
			}
		}
		maps[index]["defendant_digest"] = defendantDigests
	}*/

	return maps
}


