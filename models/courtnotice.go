package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"fmt"
)

/*type CourtNotices struct {//开庭公告表

}*/

func init()  {

	//orm.RegisterModel(new(CourtNotices))
	O = orm.NewOrm()
}

func GetCourtNoticesInfoByIds(ids []int)  {
	length := len(ids)
	where := strings.Repeat("?,",length)
	//where = where[0:-1]
	fmt.Println(where)
	//O.Raw("select * from court_notices where id in")
}

