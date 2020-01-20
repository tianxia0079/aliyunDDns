package main

import (
	"AliYunDDns/app/service/dns_jobs"
	"AliYunDDns/app/util"
	"fmt"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/gtime"
	"github.com/golang/glog"
	"time"
)

func main() {
	f3()
}
func f3() {
	_, err := gcron.Add("asdfasdf", func() {
		fmt.Println(gtime.Now())
	}, "second-cron")
	glog.Error("add", err)

	time.Sleep(3 * time.Hour)

}

func f2() {
	job := dns_jobs.QueryOneJob("6020598790754312")
	glog.Info(job)

}
func f1() {
	fmt.Println("打开浏览器")
	util.OpenUrl("http://baidu.com")
}
