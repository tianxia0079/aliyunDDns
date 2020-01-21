package main

import (
	"AliYunDDns/app/service/aliyunddns"
	"AliYunDDns/app/service/cronService"
	"AliYunDDns/app/service/dns_jobs"
	_ "AliYunDDns/boot"
	_ "AliYunDDns/router"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gsession"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang/glog"
	"time"
)

func main() {
	s := g.Server()
	s.SetServerRoot("./public")
	s.SetSessionMaxAge(6 * time.Hour)
	s.SetSessionStorage(gsession.NewStorageMemory())
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
