package main

import (
	"io/ioutil"
	"strings"
	"fmt"
)

func main(){
	//从配置文件读取要抓取的原始地址，并发线程数，轮询间隔等数据
	cmd:=ParseCmd()
	urls:=fetchUrls(cmd)
	for _,url:= range urls{
		fmt.Print("fetched url:"+url+"\n")
	}
}
func fetchUrls(cmd *Cmd) (urls []string) {
	if cmd.conf=="" {
		file,_:=ioutil.ReadFile("conf/spider.conf")
		content:=string(file)
		urls=strings.Split(content,"\n")
	}
	return urls
}