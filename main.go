package main

import (
	"AliYunDDns/app/service/aliyunddns"
	"AliYunDDns/app/service/cronService"
	"AliYunDDns/app/service/dns_jobs"
	_ "AliYunDDns/boot"
	_ "AliYunDDns/router"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang/glog"
)

func main() {
	s := g.Server()
	/*s.BindHookHandlerByMap("/*", map[string]ghttp.HandlerFunc{
		ghttp.HOOK_BEFORE_SERVE:  func(r *ghttp.Request) { glog.Println(ghttp.HOOK_BEFORE_SERVE) },
		ghttp.HOOK_AFTER_SERVE:   func(r *ghttp.Request) { glog.Println(ghttp.HOOK_AFTER_SERVE) },
		ghttp.HOOK_BEFORE_OUTPUT: func(r *ghttp.Request) { glog.Println(ghttp.HOOK_BEFORE_OUTPUT) },
		ghttp.HOOK_AFTER_OUTPUT:  func(r *ghttp.Request) { glog.Println(ghttp.HOOK_AFTER_OUTPUT) },
	})*/
	s.SetServerRoot("./public")
	s.Start()
	initConfig()
	initJobs()
	g.Wait()
}

func initJobs() {
	allToStartS := dns_jobs.AllToStart()
	for _, v := range allToStartS {
		cronService.AddCronJob(v.Cron, v.Domain)
	}
	glog.Info("从数据加载配置的任务完成.启动任务数量:" + gconv.String(len(allToStartS)) + "!")
}
func initConfig() {
	aliyunddns.InitConfig()
}
