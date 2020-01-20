package router

import (
	"AliYunDDns/app/api/help"
	"AliYunDDns/app/api/home"
	"AliYunDDns/app/api/jobs"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", home.Index)
		group.ALL("/help", help.Help)
		group.ALL("/UpdateConfig", home.UpdateConfig)
		group.ALL("/AddOrUpdateJobPage", jobs.AddOrUpdateJobPage)
		group.ALL("/AddOne", jobs.AddOne)
		group.ALL("/DeleteJob", jobs.DeleteJob)
		group.ALL("/Page", jobs.Page)
		group.ALL("/ChangeState", jobs.ChangeState)
		group.ALL("/GfJobsInfo", jobs.GfJobsInfo)

	})
}
