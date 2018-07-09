package main

import (
	"net/http"
	"errors"
	"io/ioutil"
	"regexp"
	"os"
)

func main(){
	//if checkFile("./","url.txt").size()==0 {
	//
	//}
	var url="https://my.hupu.com/search?q=%E5%AE%9E%E6%88%98%E5%88%A9%E5%99%A8"
	//下载原始页面
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic(errors.New(err.Error()))
	}
	//读取下载的页面
	bytes, _ := ioutil.ReadAll(resp.Body)
	//正则匹配需要的链接
	reg:=regexp.MustCompile("https://[a-zA-Z]+.hupu.com/*")
	file, _ := os.OpenFile("url.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	defer file.Close()
	//遍历所有匹配的url
	for _,d := range reg.FindAllString(string(bytes),-1) {
		//与url.txt中保存的url对比
		ff,_:=os.OpenFile("url.txt",os.O_RDWR,0666)
		f,_:=ioutil.ReadAll(ff)
	}
}
