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
func GetApiConfigForUrl() (key, secret, api, ip, u, p string) {
	ACCESS_KEY_ID := ""
	ACCESS_KEY_SECRET := ""
	IPAPI := ""
	Ip := ""
	Username := ""
	Password := ""

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
		if allConfig[i].Key == "UPTODATA_IP" {
			Ip = allConfig[i].Value
		}
		if allConfig[i].Key == "username" {
			Username = allConfig[i].Value
		}
		if allConfig[i].Key == "password" {
			Password = allConfig[i].Value
		}
	}
	return ACCESS_KEY_ID, ACCESS_KEY_SECRET, IPAPI, Ip, Username, Password
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
func Update(ACCESS_KEY_ID, ACCESS_KEY_SECRET, IPAPI, username, password string) string {
	db := g.DB()
	_, err1 := db.Update("dns_config", "value=?", "key='ACCESS_KEY_ID'", ACCESS_KEY_ID)
	_, err2 := db.Update("dns_config", "value=?", "key='ACCESS_KEY_SECRET'", ACCESS_KEY_SECRET)
	_, err3 := db.Update("dns_config", "value=?", "key='IPAPI'", IPAPI)

	_, err4 := db.Update("dns_config", "value=?", "key='username'", username)
	_, err5 := db.Update("dns_config", "value=?", "key='password'", password)

	if err1 == nil && err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		return "success"
	} else {
		return "false"
	}
}
func UpdateIp(ip string) string {
	db := g.DB()
	_, err := db.Update("dns_config", "value=?", "key='UPTODATA_IP'", ip)
	if err == nil {
		return "success"
	} else {
		return "false"
	}
}
func Login(username, password string) bool {
	db := g.DB()
	all, err := db.Table("dns_config").Where("key=?", "username").Where("value=?", username).
		Or("key=?", "password").Where("value=?", password).All()
	if err != nil {
		return false
	} else {
		if all != nil && len(all) == 2 {
			return true
		} else {
			return false
		}
	}
}
