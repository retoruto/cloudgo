package service

import (
	"github.com/go-martini/martini"  /*使用martini*/
)

func NewServer(port string) {   /*新建服务器*/
	m := martini.Classic()
	/*添加参数[name]martini的参数中*/
	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"] + "でていけ!"/*你好世界，快滚*/
	})
	/*对应mian函数中的端口*/
	m.RunOnAddr(":"+port)	
}
/*test ip*/
/*http://localhost:9090/hello/world*/
/*test curl*/
/*curl -v http://localhost:9090/*/
/*test ab*/
/*ab -n 1000 -c 100 http://localhost:9090/hello/world*/