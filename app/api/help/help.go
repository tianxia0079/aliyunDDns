package help

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Help(r *ghttp.Request) {
	r.Response.WriteTpl("help.html", g.Map{"v": "1"})
}
