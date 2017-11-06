# 服务计算作业4——cloudgo（基本要求）

花了一个下午才搞懂这次实验的要求是什么。参考了大佬的README.md之后才勉强完成了基本要求。


-------------------
##**过程**：

1. 按照老师给出的代码实现一遍流程：
老师的代码：https://github.com/pmlpml/golang-learning
下载zip后解压。
我这里的做法可能有点蠢：
![这里写图片描述](http://img.blog.csdn.net/20171106121806145?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

找到service文件
![这里写图片描述](http://img.blog.csdn.net/20171106121827213?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

放在GOPATH的src文件中
![这里写图片描述](http://img.blog.csdn.net/20171106121851863?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

修改main函数中service包的地址
![这里写图片描述](http://img.blog.csdn.net/20171106121919583?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

运行代码：
![这里写图片描述](http://img.blog.csdn.net/20171106121959095?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
现在就开始监听了。
打开另一个终端输入老师提供的地址即可完成老师提供的实验，这里先不截图了。

2.仿照进行实验：
main.go函数几乎不变：

```
package main

import (
    "os"
    "service"
    flag "github.com/spf13/pflag"
)

const (
    PORT string = "8080" /*设置默认的端口为8080*/
)

func main() {

    port := os.Getenv("PORT") /*如果没有监听到端口，那么端口为8080*/
    if len(port) == 0 {
        port = PORT
    }
    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")/*设置端口*/
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
	}
	
    service.NewServer(port)/*启动服务器*/
}
```
service函数：
（要放在service包中）

```
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
/*curl -v http://localhost:9090/hello/world*/
/*test ab*/
/*ab -n 1000 -c 100 http://localhost:9090/hello/world*/
```

我这里使用的是martini框架，也是看了大佬的完成之后才选的。
因为它看起来比较简单【。。。】
可能在编译main.go的时候，会出现报错缺少包：go-martini/martini
这个时候不需要慌张，打开浏览器输入：“github.com/go-martini/martini”
下载zip：
解压放在你的
![这里写图片描述](http://img.blog.csdn.net/20171106124318851?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
里面是这样的：
![这里写图片描述](http://img.blog.csdn.net/20171106124421843?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
就是把上面的zip压缩文件解压，然后修改一下文件名变成这样就可以编译了。
出现缺少的github.com/xx文件就去github找文件就好了。
因为我虚拟机连不上网，只能用这个方法了。

3.测试：
找到main函数并运行：
![这里写图片描述](http://img.blog.csdn.net/20171106124721005?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


test ip：
打开浏览器输入：
![这里写图片描述](http://img.blog.csdn.net/20171106124903812?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
可以看到监听到了网页请求的信息。

test curl
打开另一个终端输入：
`curl -v http://localhost:9090/hello/world`
![这里写图片描述](http://img.blog.csdn.net/20171106125348811?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
也可以监听到了。

test ab
打开另一个终端输入：
`ab -n 1000 -c 100 http://localhost:9090/hello/world`
![这里写图片描述](http://img.blog.csdn.net/20171106125637380?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
![这里写图片描述](http://img.blog.csdn.net/20171106125650973?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
经过压力测试，可以看到发送了1000个请求。
每一个请求花费时间为0.5ms
一半的请求需要47ms
全部请求花费了92ms