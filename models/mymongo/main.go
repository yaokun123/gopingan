package mymongo

import (
	"github.com/globalsign/mgo"
	"github.com/astaxie/beego"
	"github.com/globalsign/mgo/bson"
	"fmt"
)

var MgoSession *mgo.Session


func init()  {
	host := beego.AppConfig.String("mongo_host")
	database := beego.AppConfig.String("mongo_db")
	username := beego.AppConfig.String("mongo_user")
	passwd := beego.AppConfig.String("mongo_passwd")

	diaInfo := &mgo.DialInfo{
		Addrs: []string{host},
		Direct:    false,
		PoolLimit: 4096,
		Database:database,
		Username:username,
		Password:passwd,
	}

	var err error
	MgoSession,err = mgo.DialWithInfo(diaInfo)
	if err != nil{
		fmt.Println("======",err)
	}
}

type companyname struct {
	CompanyName string `bson:"company_name"`
}

func GetMongoInfoByDigest(digest string){
	var aObj companyname
	query := func(c *mgo.Collection) error{//匿名函数
		return c.Find(bson.M{"company_name_digest":digest}).One(&aObj)
	}
	err := handlerCollection("ic",query)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(aObj)
}

type dbCollection func(collection *mgo.Collection) (err error)

func  handlerCollection(colName string,query dbCollection) error {
	s := MgoSession.Copy()
	defer s.Close()
	c := s.DB("utn_ic").C(colName)
	return query(c)
}

