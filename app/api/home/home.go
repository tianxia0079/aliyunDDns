package home

import (
	dns_configserver "AliYunDDns/app/service/dns_config_service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	_ "github.com/mattn/go-sqlite3"
)

func Index(r *ghttp.Request) {
	errorinfo := ""

	key, secret, api := dns_configserver.GetApiConfigForUrl()

	r.Response.WriteTpl("home.html", g.Map{
		"errorinfo":         errorinfo,
		"configs":           nil,
		"ACCESS_KEY_ID":     key,
		"ACCESS_KEY_SECRET": secret,
		"IPAPI":             api,
	})
}
func UpdateConfig(r *ghttp.Request) {
	//glog.Debug("参数:", r.Form, r.GetString("ACCESS_KEY_ID"))

	ACCESS_KEY_ID := r.GetString("ACCESS_KEY_ID")
	ACCESS_KEY_SECRET := r.GetString("ACCESS_KEY_SECRET")
	IPAPI := r.GetString("IPAPI")

	status := dns_configserver.Update(ACCESS_KEY_ID, ACCESS_KEY_SECRET, IPAPI)
	r.Response.Write(status)
}
