package cronService

import (
	"AliYunDDns/app/service/aliyunddns"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/util/guuid"
)

func RunOne(state int, domain string) (code, message string) {
	if state == 0 {
		aliyunddns.UpdateAllTypeA("[手动触发模式]", domain)
		return "success", "手动触发执行成功"
	} else if state == 1 {
		//运行状态 触发一次即可
		gcron.Start(domain)
		return "success", "手动触发执行成功"
	} else {
		return "fail", "未知运行状态"
	}
}

//更新已经运行的任务
func UpdateRunCronJob(cron, domain string) {
	RemoveCronJob(domain)
	AddCronJob(cron, domain)
}

/**
定时任务添加一个协程阻塞了，会不会阻塞web主线程测试
*/
func AddCronJob(cron, domain string) {
	gcron.AddSingleton(cron, func() {
		aliyunddns.UpdateAllTypeA(cron, domain)
	}, domain) //通过设置定时任务名控制任务状态
	//todo:目前用domain管理即可，因为一个domain下只有一个任务,后续扩展出新场景，用id管理
}
func RemoveCronJob(domain string) {
	search := gcron.Search(domain)
	if search != nil && search.Name == domain {
		//确认找到 再删除
		gcron.Remove(domain)
	}
}

//todo:校验cron  待gf有cron校验方法，替换
func ValidateCron(cron string) bool {
	uuid, _ := guuid.NewUUID()
	uuidName := uuid.String()
	_, err := gcron.AddSingleton(cron, func() {
		//仅测试cron是否合法
	}, uuidName)
	if err != nil {
		return false
	}
	gcron.Remove(uuidName)
	return true
}
