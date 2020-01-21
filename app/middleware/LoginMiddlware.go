package middleware

import (
	"github.com/gogf/gf/net/ghttp"
)

/**
拦截根路径的中间件
*/
const SESSIONKEY = "UserKey_SESSION"

var URLFILTER []string = []string{"/", "/Login"}

func MiddlewareAuth(r *ghttp.Request) {
	next := false
	//url相对路径
	path := r.Request.URL.Path
Loop:
	for _, v := range URLFILTER {
		if v == path {
			next = true
			break Loop
		}
	}

	if next {
		//非拦截请求
		r.Middleware.Next()
	} else {
		if r.Session.GetString(SESSIONKEY) == "" {
			r.Response.RedirectTo("/Login", 302)
		} else {
			//已登录
			r.Middleware.Next()
		}
	}

}
