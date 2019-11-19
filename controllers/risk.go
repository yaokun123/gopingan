package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"gopingan/models"
)

type RiskController struct {
	beego.Controller
}

func (this *RiskController) GetInfo()  {
	digest := this.GetString("digest")
	fmt.Println(digest)

	company_name := "小米科技有限责任公司"
	result,allCount := models.GetInfoByName(company_name,0)
	//var idList []int

	for item := range result{
		//append(idList, item)
		fmt.Println(item)
	}
	fmt.Println(allCount)
}
