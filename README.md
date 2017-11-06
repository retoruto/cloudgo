# 服务计算作业4——cloudgo（基本要求）

花了一个下午才搞懂这次实验的要求是什么。参考了大佬的README.md之后才勉强完成了基本要求。


-------------------
**过程**：

我这里使用的是martini框架，它的教程很详细，比较友好，而且相比其他框架容易上手</br>
可能在编译main.go的时候，会出现报错缺少包：go-martini/martini
这个时候不需要慌张，打开浏览器输入：“github.com/go-martini/martini”
下载zip：
解压放在你的
![这里写图片描述](http://img.blog.csdn.net/20171106124318851?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
</br>
里面是这样的：
![这里写图片描述](http://img.blog.csdn.net/20171106124421843?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
</br>
就是把上面的zip压缩文件解压，然后修改一下文件名变成这样就可以编译了。</br>
出现缺少的github.com/xx文件就去github找文件就好了。
因为我虚拟机连不上网，只能用这个方法了。
</br>
3.测试：
找到main函数并运行：
![这里写图片描述](http://img.blog.csdn.net/20171106124721005?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

</br>
test ip：
打开浏览器输入：
![这里写图片描述](http://img.blog.csdn.net/20171106124903812?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
</br>
可以看到监听到了网页请求的信息。
</br>
test curl
打开另一个终端输入：
`curl -v http://localhost:9090/hello/world`
![这里写图片描述](http://img.blog.csdn.net/20171106125348811?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
也可以监听到了。
</br>
test ab
打开另一个终端输入：
`ab -n 1000 -c 100 http://localhost:9090/hello/world`
![这里写图片描述](http://img.blog.csdn.net/20171106125637380?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
![这里写图片描述](http://img.blog.csdn.net/20171106125650973?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
经过压力测试，可以看到发送了1000个请求。</br>
每一个请求花费时间为0.5ms</br>
一半的请求需要47ms</br>
全部请求花费了92ms</br>
