package home

import (
	"AliYunDDns/app/middleware"
	dns_configserver "AliYunDDns/app/service/dns_config_service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/golang/glog"
	_ "github.com/mattn/go-sqlite3"
)

type LoginUser struct {
	Username string `p:"username"`
	Password string `p:"password"`
}

func Login(r *ghttp.Request) {
	method := r.Request.Method
	if method == "GET" {
		r.Response.WriteTpl("login.html", g.Map{
			"info": "",
		})
	} else {
		var user *LoginUser
		if err := r.Parse(&user); err == nil {
			glog.Info("参数:", user)
			login := dns_configserver.Login(user.Username, user.Password)
			if login {
				r.Session.Set(middleware.SESSIONKEY, user.Username)
			}
			r.Response.Write(login)
		} else {
			r.Response.Write("error:" + err.Error())
		}
	}

}
func Main(r *ghttp.Request) {
	errorinfo := ""

	key, secret, api, ip, username, password := dns_configserver.GetApiConfigForUrl()

	r.Response.WriteTpl("home.html", g.Map{
		"errorinfo":         errorinfo,
		"configs":           nil,
		"ACCESS_KEY_ID":     key,
		"ACCESS_KEY_SECRET": secret,
		"IPAPI":             api,
		"UPTODATA_IP":       ip,
		"username":          username,
		"password":          password,
	})
}
func UpdateConfig(r *ghttp.Request) {
	//glog.Debug("参数:", r.Form, r.GetString("ACCESS_KEY_ID"))

	ACCESS_KEY_ID := r.GetString("ACCESS_KEY_ID")
	ACCESS_KEY_SECRET := r.GetString("ACCESS_KEY_SECRET")
	IPAPI := r.GetString("IPAPI")
	username := r.GetString("username")
	password := r.GetString("password")

	status := dns_configserver.Update(ACCESS_KEY_ID, ACCESS_KEY_SECRET, IPAPI, username, password)
	r.Response.Write(status)
}
