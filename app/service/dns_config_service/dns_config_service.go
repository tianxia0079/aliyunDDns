package dns_config

import (
	"github.com/gogf/gf/frame/g"
	"github.com/golang/glog"
)

type Entity struct {
	Id    int    `orm:"id,primary" json:"id"`    //
	Key   string `orm:"key"    json:"key"`       //
	Value string `orm:"value"      json:"value"` //
	Time  string `orm:"time"    json:"time"`     //
}

func QueryConfig1() {
	db := g.DB()
	all, _ := db.Table("dns_config").FindAll()
	glog.Info(all)
}
func GetApiConfigForUrl() (key, secret, api string) {
	ACCESS_KEY_ID := ""
	ACCESS_KEY_SECRET := ""
	IPAPI := ""
	allConfig := QueryConfig2()
	for i := range allConfig {
		if allConfig[i].Key == "ACCESS_KEY_ID" {
			ACCESS_KEY_ID = allConfig[i].Value
		}
		if allConfig[i].Key == "ACCESS_KEY_SECRET" {
			ACCESS_KEY_SECRET = allConfig[i].Value
		}
		if allConfig[i].Key == "IPAPI" {
			IPAPI = allConfig[i].Value
		}
	}
	return ACCESS_KEY_ID, ACCESS_KEY_SECRET, IPAPI
}
func QueryConfig2() []*Entity {
	var list []*Entity
	db := g.DB()
	all, _ := db.Table("dns_config").FindAll()
	//glog.Info(all)
	err := all.Structs(&list)
	if err != nil {
		glog.Error(err)
	}
	return list
}
func Update(ACCESS_KEY_ID, ACCESS_KEY_SECRET, IPAPI string) string {
	db := g.DB()
	_, err1 := db.Update("dns_config", "value=?", "key='ACCESS_KEY_ID'", ACCESS_KEY_ID)
	_, err2 := db.Update("dns_config", "value=?", "key='ACCESS_KEY_SECRET'", ACCESS_KEY_SECRET)
	_, err3 := db.Update("dns_config", "value=?", "key='IPAPI'", IPAPI)
	if err1 == nil && err2 == nil && err3 == nil {
		return "success"
	} else {
		return "false"
	}
}
