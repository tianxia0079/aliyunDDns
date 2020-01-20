package jobs

import (
	"AliYunDDns/app/service/cronService"
	"AliYunDDns/app/service/dns_jobs"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang/glog"
)

type ResultInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

//通过为struct属性绑定
type LayuiPage struct {
	Page  int `p:'page' json:"page"`
	Limit int `p:'limit' json:"limit"`
}
type LayiUiTable struct {
	Code  int                `json:"code"`
	Msg   string             `json:"msg"`
	Count int                `json:"count"`
	Data  []*dns_jobs.Entity `json:"data"`
}

func RunOne(r *ghttp.Request) {
	state := r.GetInt("state")
	domain := r.GetString("domain")

	code, message := cronService.RunOne(state, domain)
	re := &ResultInfo{
		Code:    code,
		Message: message,
	}
	r.Response.WriteJson(re)
}
func ChangeState(r *ghttp.Request) {
	id := r.GetString("id")
	state := r.GetInt("state")
	domain := r.GetString("domain")
	cron := r.GetString("cron")

	if state == 0 {
		state = 1
	} else if state == 1 {
		state = 0
	}
	status := dns_jobs.UpdateState(id, state)
	if status == "success" {
		if state == 0 {
			//需要更新为0 关闭
			cronService.RemoveCronJob(domain)
		} else if state == 1 {
			//需要更新为1 启动
			cronService.AddCronJob(cron, domain)
		}
	}
	r.Response.Write(status)
}

func Page(r *ghttp.Request) {
	var page *LayuiPage

	if err := r.Parse(&page); err != nil {
		glog.Error("分页参数转struct异常:", err)
	}

	jobs, count := dns_jobs.QueryJobs(page.Page, page.Limit)

	t := LayiUiTable{Count: count, Msg: "", Code: 0, Data: jobs}
	r.Response.Write(t)
}
func DeleteJob(r *ghttp.Request) {
	id := r.GetString("id")
	domain := r.GetString("domain")
	status := dns_jobs.DeleteOneJob(id)
	if status == "success" {
		cronService.RemoveCronJob(domain)
	}
	r.Response.Write(status)
}
func GfJobsInfo(r *ghttp.Request) {
	entries := gcron.Entries()
	glog.Info("所有任务：", &entries)
	r.Response.WriteJson(entries)
}
func AddOne(r *ghttp.Request) {
	id := r.GetString("id")
	comment := r.GetString("comment")
	domain := r.GetString("domain")
	cron := r.GetString("cron")
	state := r.GetInt("state")

	resultInfo := ResultInfo{}

	validateCron := cronService.ValidateCron(cron)
	if !validateCron {
		//cron不合法
		resultInfo.Code = "fail"
		resultInfo.Message = "cron表达式不合法!"
		//结束后续业务逻辑
		r.Response.Write(resultInfo)
	} else {
		if id == "" {
			//add
			status, err := dns_jobs.AddJob(domain, comment, cron)
			resultInfo.Code = status
			resultInfo.Message = err
			//add后不自动启动
			//只有add时候才需要自动启动
			//cornService.AddCronJob(cron, domain)
		} else {
			//edit
			status := dns_jobs.EditJob(id, domain, comment, cron)
			resultInfo.Code = status
			//如果编辑的任务已经启动 要更新cron
			if state == 1 {
				glog.Info("该任务已经运行中 重新加载参数")
				cronService.UpdateRunCronJob(cron, domain)
			}
		}
		r.Response.Write(resultInfo)
	}

}
func AddOrUpdateJobPage(r *ghttp.Request) {
	opeType := ""
	var oneJob *dns_jobs.Entity
	jobid := r.GetQueryString("id")

	if jobid == "" {
		//add
		opeType = "add"
	} else {
		//edit
		opeType = "edit"
		oneJob = dns_jobs.QueryOneJob(jobid)
	}
	//glog.Info("参数 返回值:", jobid, oneJob)
	r.Response.WriteTpl("jobManager.html", gconv.Map(oneJob), g.Map{
		"type": opeType,
	})
}
