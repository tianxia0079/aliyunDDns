package aliyunddns

import (
	dns_configserver "AliYunDDns/app/service/dns_config_service"
	"github.com/gogf/gf/net/ghttp"
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
*/
func DDnsUpdateAllTypeA(domain string, wait *sync.WaitGroup) {
	defer wait.Done()

	result, err := ghttp.Get(api)
	defer result.Body.Close()

	if err != nil {
		glog.Error("获取公网ip接口异常", err)
		return
	}

	ip, _ := ioutil.ReadAll(result.Body)
	glog.Info("当前最新ip:" + string(ip))
	lis := List(domain)
	updateList := []alidns.Record{}
	for _, v := range lis {
		if v.Type == "A" {
			updateList = append(updateList, v)
		}
	}
	glog.Info("需要更新的a类型：", updateList)
	for _, v := range updateList {
		//如果异常不处理，就没必要让当前协程等待这个匿名协程结果
		glog.Info("现在更新子域名:", v.RR)
		client, _ := getClient()
		request := alidns.CreateUpdateDomainRecordRequest()
		request.RecordId = v.RecordId
		request.RR = v.RR
		request.Type = v.Type
		request.Value = string(ip)
		_, err := client.UpdateDomainRecord(request)
		if err != nil {
			if gstr.Contains(err.Error(), "DomainRecordDuplicate") {
				glog.Info("		当前公网ip和远程A记录Ip一样,无须更新!")
			} else {
				//todo:异常处理
				glog.Error("	调用更新接口时候未知异常:", err.Error())
			}
		} else {
			glog.Info("		检测到IP变化,更新IP成功!")
		}
	}
}
