package router

import (
	"AliYunDDns/app/api/help"
	"AliYunDDns/app/api/home"
	"AliYunDDns/app/api/jobs"
	"AliYunDDns/app/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.MiddlewareAuth)
		group.ALL("/", home.Login)
		group.ALL("/Login", home.Login)
		group.ALL("/Main", home.Main)
		group.ALL("/Help", help.Help)
		group.ALL("/UpdateConfig", home.UpdateConfig)
		group.ALL("/AddOrUpdateJobPage", jobs.AddOrUpdateJobPage)
		group.ALL("/AddOne", jobs.AddOne)
		group.ALL("/DeleteJob", jobs.DeleteJob)
		group.ALL("/Page", jobs.Page)
		group.ALL("/ChangeState", jobs.ChangeState)
		group.ALL("/GfJobsInfo", jobs.GfJobsInfo)
		group.ALL("/RunOne", jobs.RunOne)

	})
}
