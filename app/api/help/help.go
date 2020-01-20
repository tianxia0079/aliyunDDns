package help

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Help(r *ghttp.Request) {
	/*go func() {
		time.Sleep(1 * time.Hour)
	}()*/

	r.Response.WriteTpl("help.html", g.Map{"v": "1"})
}
