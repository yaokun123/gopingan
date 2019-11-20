package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"gopingan/models/mymongo"
	"gopingan/models"
)

type RiskController struct {
	beego.Controller
}

func (this *RiskController) GetInfo()  {
	digest := this.GetString("digest")
	company_name := mymongo.GetMongoInfoByDigest(digest)

	//开庭公告
	courtNoticeMapResult,courtNoticeMapCount := models.GetCourtNoticesMapInfoByName(company_name,0)
	var idList []int
	for _,item := range courtNoticeMapResult{
		idList = append(idList, item.Id)
	}
	courtNoticeResult := models.GetCourtNoticesInfoByIds(idList)

	fmt.Println(courtNoticeMapCount)
	fmt.Println(courtNoticeResult)
	this.Data["json"] = courtNoticeResult
	this.ServeJSON()
}