package cronService

import (
	"AliYunDDns/app/service/aliyunddns"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/guuid"
	"github.com/golang/glog"
	"sync"
)

//更新已经运行的任务
func UpdateRunCronJob(cron, domain string) {
	RemoveCronJob(domain)
	AddCronJob(cron, domain)
}
func AddCronJob(cron, domain string) {
	gcron.AddSingleton(cron, func() {
		wait := sync.WaitGroup{}
		wait.Add(1)
		go aliyunddns.DDnsUpdateAllTypeA(domain, &wait)
		wait.Wait()
		glog.Info(gtime.Now(), " domain:", domain, " cron: ", cron, " 更新一次!")
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
