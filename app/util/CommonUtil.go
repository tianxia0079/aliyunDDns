package util

// 打开系统默认浏览器

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gproc"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/golang/glog"
	"os/exec"
	"runtime"
)

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func OpenUrl(uri string) {
	system := runtime.GOOS
	switch system {
	case "windows":
		{
			// 无GUI调用
			exec.Command(`cmd`, `/c`, `start`, uri).Start()
			break
		}
	case "linux":
		{
			exec.Command(`xdg-open`, `https://www.jianshu.com`).Start()
			break
		}
	case "darwin":
		{
			exec.Command(`open`, `https://www.jianshu.com`).Start()
			break
		}
	}
}

//2020-01-01 09:11:27
func GetbeijingTime() string {
	re := ""
	content := ghttp.GetContent("http://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp")
	taobaoTime := new(TaobaoTime)
	err := gjson.DecodeTo(content, &taobaoTime)
	if err != nil {
		glog.Error(err)
		return ""
	}
	timestamptemp := (gconv.Int64(taobaoTime.Data.T)) / 1000
	re = gtime.NewFromTimeStamp(timestamptemp).String()
	glog.Info(re)
	return re
}
func UpdateSystemDateAuto() {
	UpdateSystemDate(GetbeijingTime())
}

//2020-1-1 16:55:51 格式
func UpdateSystemDate(dateTime string) bool {
	system := runtime.GOOS
	switch system {
	case "windows":
		{
			_, err1 := gproc.ShellExec(`date  ` + gstr.Split(dateTime, " ")[0])
			_, err2 := gproc.ShellExec(`time  ` + gstr.Split(dateTime, " ")[1])
			if err1 != nil && err2 != nil {
				glog.Info("更新系统时间错误:请用管理员身份启动程序!")
				return false
			}
			glog.Info("已自动同步时间: ", dateTime)
			return true
			break
		}
	case "linux":
		{
			_, err1 := gproc.ShellExec(`date -s  "` + dateTime + `"`)
			if err1 != nil {
				glog.Info("更新系统时间错误:", err1.Error())
				return false
			}
			glog.Info("已自动同步时间: ", dateTime)
			return true
			break
		}
	case "darwin":
		{
			//todo:mac是否可以执行 未测试
			_, err1 := gproc.ShellExec(`date -s  "` + dateTime + `"`)
			if err1 != nil {
				glog.Info("更新系统时间错误:", err1.Error())
				return false
			}
			glog.Info("已自动同步时间: ", dateTime)
			return true
			break
		}
	}
	return false
}

type TimeStamp struct {
	T string `json:t`
}
type TaobaoTime struct {
	Api  string    `json:api`
	V    string    `json:v`
	ret  string    `json:ret`
	Data TimeStamp `json:data`
}
