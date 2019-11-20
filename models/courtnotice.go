package models

import (
	"strings"
	_"gopingan/models/dbresource/utnngrisk"
	"github.com/astaxie/beego/orm"
	"fmt"
	"unicode/utf8"
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

	//公司名检测plaintiff/defendant
	var company_name_list []string
	for index,item := range maps {
		if item["startDate"] != nil{
			maps[index]["startDate"] = item["startDate"].(string)[0:10]
		}
		if item["update_date"] != nil{
			maps[index]["update_date"] = item["update_date"].(string)[0:10]
		}
		if item["create_time"] != nil{
			maps[index]["create_time"] = item["create_time"].(string)[0:10]
		}
		plaintiffs := strings.Split(item["plaintiff"].(string),",")
		for index2,company_name := range plaintiffs{
			if utf8.RuneCountInString(company_name) > 4{//小于4个字的不取公司名
				company_name_list = append(company_name_list,company_name)
			}else if utf8.RuneCountInString(company_name) == 3{
				tmp := []rune(company_name)
				first_char := string(tmp[0:1])
				end_char := string(tmp[2:1])
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

		defendants := strings.Split(item["defendant"].(string),",")
		for index3,company_name := range defendants{
			if utf8.RuneCountInString(company_name) > 4{//小于4个字的不取公司名
				company_name_list = append(company_name_list,company_name)
			}else if utf8.RuneCountInString(company_name) == 3{
				tmp := []rune(company_name)
				first_char := string(tmp[0:1])
				end_char := string(tmp[2:1])
				result_name := first_char + "*" + end_char
				plaintiffs[index3] = result_name
			}else if utf8.RuneCountInString(company_name) == 2{
				tmp := []rune(company_name)
				first_char := string(tmp[0:1])
				result_name := first_char + "*"
				plaintiffs[index3] = result_name
			}
		}
		maps[index]["defendant"] = defendants
	}

	fmt.Println(maps)
}


