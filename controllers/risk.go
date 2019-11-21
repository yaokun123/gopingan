package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"gopingan/models/mymongo"
	"gopingan/models"
	"github.com/astaxie/beego/orm"
	"sync"
)

type RiskController struct {
	beego.Controller
}

func (this *RiskController) GetInfo()  {
	digest := this.GetString("digest")
	company_name := mymongo.GetMongoInfoByDigest(digest)

	//创建waitGroup
	var wg sync.WaitGroup


	//1、开庭公告
	courtNoticeResultChannel := make(chan []orm.Params,2)
	courtNoticeMapCountChannel := make(chan int64,2)
	wg.Add(1)
	go getCourtNotice(company_name,courtNoticeResultChannel,courtNoticeMapCountChannel,&wg)


	//2、法律诉讼
	companyLawsuitParsedInfoResultChannel := make(chan []orm.Params,2)
	companyLawsuitParsedInfoMapCountChannel := make(chan int64,2)
	wg.Add(1)
	go getCompanyLawsuitParsedInfo(company_name,companyLawsuitParsedInfoResultChannel,companyLawsuitParsedInfoMapCountChannel,&wg,digest)


	//数据统一在后面取
	wg.Wait()


	//开庭公告的数据
	courtNoticeResult := <- courtNoticeResultChannel
	courtNoticeMapCount := <- courtNoticeMapCountChannel
	fmt.Println(courtNoticeMapCount)

	//法律诉讼的权限
	companyLawsuitParsedInfoResult := <- companyLawsuitParsedInfoResultChannel
	companyLawsuitParsedInfoMapCount := <- companyLawsuitParsedInfoMapCountChannel
	fmt.Println(companyLawsuitParsedInfoMapCount)


	//返回数据
	result := make(map[string]interface{})
	result["courtNotice"] = courtNoticeResult//开庭公告
	result["companyLawsuitParsedInfo"] = companyLawsuitParsedInfoResult//法律诉讼
	this.Data["json"] = result
	this.ServeJSON()
}


//开庭公告
func getCourtNotice(company_name string,courtNoticeResultChannel chan []orm.Params,courtNoticeMapCountChannel chan int64,wg *sync.WaitGroup) {
	courtNoticeMapResult,courtNoticeMapCount := models.GetCourtNoticesMapInfoByName(company_name,0)
	var idList []int
	for _,item := range courtNoticeMapResult{
		idList = append(idList, item.Id)
	}
	courtNoticeResult := models.GetCourtNoticesInfoByIds(idList)

	courtNoticeResultChannel <- courtNoticeResult
	courtNoticeMapCountChannel <- courtNoticeMapCount
	wg.Done()
}

//法律诉讼
func getCompanyLawsuitParsedInfo(company_name string,companyLawsuitParsedInfoResultChannel chan []orm.Params,companyLawsuitParsedInfoMapCountChannel chan int64,wg *sync.WaitGroup,digest string){
	companyLawsuitParsedInfoMapResult,companyLawsuitParsedInfoMapCount := models.GetCompanyLawsuitParsedInfoMapInfoByName(company_name,0)
	var uuidList []string
	for _,item := range companyLawsuitParsedInfoMapResult{
		uuidList = append(uuidList,item.Uuid)
	}
	companyLawsuitParsedInfoResult := models.GetCompanyLawsuitParsedInfoByUuids(uuidList,digest)

	companyLawsuitParsedInfoResultChannel <- companyLawsuitParsedInfoResult
	companyLawsuitParsedInfoMapCountChannel <- companyLawsuitParsedInfoMapCount
	wg.Done()
}