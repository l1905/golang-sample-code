package main

import (
	"flag"
	"fmt"
	"qing_hospital/guahao"
	"strings"
)

var (
	date_zone string
	time string
)

func init() {
	flag.StringVar(&date_zone, "date_zone", "", "设置时间区间，比如 2019-05-06,2019-05-07， 不传默认全部")
	flag.StringVar(&time, "time", "", "设置区间， 比如是希望获取上午，还是下午， 传参数 比如 PM, AM, 不传默认全部")
}


func main() {

	//需要提前初始化手机号，+ 钉钉token + 微信token

	flag.Parse()

	//打印输出 参数
	fmt.Println(date_zone)
	fmt.Println(time)

	//请求URL 获取返回值
	body_str := guahao.RquestChanke()
	//正则解析 获取结构内容，并且代码
	time = strings.ToUpper(time)
	print_text := guahao.Regex(body_str, date_zone, time)

	fmt.Println(print_text)
	//发送报警
	if len(print_text) > 0 {
		//发送报警 微信 && 钉钉
		guahao.SendAlarm(print_text)
	} else {
		fmt.Println("暂时无数据")
	}

}