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
	courtNoticeResultChannel := make(chan []orm.Params)
	courtNoticeMapCountChannel := make(chan int64)
	wg.Add(1)
	go getCourtNotice(company_name,courtNoticeResultChannel,courtNoticeMapCountChannel,&wg)


	fmt.Println("haha")
	wg.Wait()
	fmt.Println("哈哈")
	courtNoticeResult := <- courtNoticeResultChannel
	courtNoticeMapCount := <- courtNoticeMapCountChannel
	fmt.Println(courtNoticeMapCount)
	this.Data["json"] = courtNoticeResult
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