package domains

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

var VERSION string = ""
var IPAPI string = ""
var DOMAINS []interface{} = nil

func init() {
	UpdateVars()
}
func UpdateVars() {
	if DOMAINS == nil {
		DOMAINS = g.Cfg().GetArray("aliyun.DOMAINS")
		glog.Debug("读取了一次配置文件")
	}
	if VERSION == "" {
		VERSION = g.Cfg().GetString("aliyun.VERSION")
	}
	if IPAPI == "" {
		IPAPI = g.Cfg().GetString("aliyun.IPAPI")
	}
}

//顶级域名查询列表
func DomainList() (arr []interface{}) {
	return DOMAINS
}
func Version() (str string) {
	return VERSION
}
func Ipapi() (str string) {
	return IPAPI
}
