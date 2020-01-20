package aliyunddns

import (
	dns_configserver "AliYunDDns/app/service/dns_config_service"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"io/ioutil"
)

/**
接口地址
https://help.aliyun.com/document_detail/29774.html?spm=a2c4g.11186623.6.653.3bdb5eb4Ou50mX
*/
import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/gogf/gf/os/glog"
	"sync"
)

var key = ""
var secret = ""
var api = ""

//系统启动加载参数；更新配置更新参数
func InitConfig() {
	key, secret, api = dns_configserver.GetApiConfigForUrl()
}
func getClient() (client *alidns.Client, err error) {
	return alidns.NewClientWithAccessKey("", key, secret)
}
func List(domain string) []alidns.Record {
	client, _ := getClient()
	request := alidns.CreateDescribeDomainRecordsRequest()
	request.Scheme = "https"
	request.DomainName = domain
	response, err := client.DescribeDomainRecords(request)
	if err != nil {
		glog.Error("获取"+domain+"子域名失败:", err.Error())
	}
	res := response.DomainRecords
	//glog.Debug("当前域名下的子域名:" + gconv.String(res))
	return res.Record
}

/**
特殊业务场景定义：
把所有A记录ip更新
todo:测试定时任务会不会阻塞web主线程
*/
func UpdateAllTypeA(cron, domain string) string {
	defer glog.Info("执行A记录更新ip...结束")
	re := "执行A记录更新ip...开始\n"

	re += ("执行时间: " + gtime.Now().String() + "\n")
	glog.Info(gtime.Now(), " domain:", domain, " cron: ", cron, " 执行了!")
	glog.Info("开始请求api")
	result, err := ghttp.Get(api)
	glog.Info("结束请求api")

	defer result.Body.Close()

	if err != nil {
		re = "获取公网ip接口异常:" + err.Error()
	} else {
		ip, _ := ioutil.ReadAll(result.Body)
		//glog.Info("		当前公网 ip:" + string(ip))
		re += "当前公网 ip:" + string(ip) + "\n"
		list := List(domain)
		updateList := []alidns.Record{}
		for _, v := range list {
			if v.Type == "A" {
				updateList = append(updateList, v)
			}
		}
		var wait = sync.WaitGroup{}
		wait.Add(len(updateList))
		infoback := make(chan string, len(updateList))

		for _, val := range updateList {
			go func(info chan string, v alidns.Record, w *sync.WaitGroup) {
				defer w.Done()
				re := "		现在更新子域名:" + v.RR
				client, _ := getClient()
				request := alidns.CreateUpdateDomainRecordRequest()
				request.RecordId = v.RecordId
				request.RR = v.RR
				request.Type = v.Type
				request.Value = string(ip)
				_, err := client.UpdateDomainRecord(request)
				if err != nil {
					if gstr.Contains(err.Error(), "DomainRecordDuplicate") {
						re += " 当前公网ip和远程A记录Ip一样,无须更新!"
					} else {
						re += " 调用更新接口时候未知异常:" + err.Error()
					}
				} else {
					re += "		检测到IP变化,更新IP成功!"
				}
				re += "\n"
				//把任务执行情况写入通信chal info里
				info <- re
			}(infoback, val, &wait)
		}
		wait.Wait()
		//配合wait 此时可以关闭chan
		//必须是带指定缓存的chan才可以这样读写chan
		close(infoback)

		//关闭之后的chan可以多种安全方式读取
		//method1 直接for循环取
		for v := range infoback {
			re += v
		}
	}
	glog.Info(re)
	return re
}
