package guahao

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func RquestChanke() (body_str string)  {
	url := "https://api.yuantutech.com/user-web/restapi/common/reservation/listScheduleinfoByDate?corpId=261&deptCode=0080&regMode=1&unionId=29&invokerChannel=H5&invokerDeviceType=weixin&invokerAppVersion=2.1.0&t=2067&callback=jsonp1555727870452"

	//resp, err := http.Get(url)
	//if err != nil {
	//	// handle error
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//
	//fmt.Println(string(body))
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("返回错误")
	}

	//还可以这样设置header， 另一种设置
	//req.Headers = map[string]string{
	//	"Accept-Encoding": "",
	//}

	req.Header.Add("Host", "api.yuantutech.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 7.1.1; MI MAX 2 Build/NMF26F; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/044611 Mobile Safari/537.36 MMWEBID/9053 MicroMessenger/7.0.3.1400(0x2700033C) Process/tools NetType/WIFI Language/zh_CN`)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Referer", "https://s.yuantutech.com/yuantu/h5-cli/2.1.0/appointment-select.html?deptCode=0080&corpId=261&deptName=%E4%BA%A7%E7%A7%91%E9%97%A8%E8%AF%8A&parentDeptCode=&regMode=1&regType=1&unionId=29&target=_blank")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br") //需要注释掉， 否则乱码
	req.Header.Add("Accept-Language", "zh-CN,en-US;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("返回错误")
	}
	body, err := ioutil.ReadAll(resp.Body)

	//body_str := string(body)
	//fmt.Println(string(body))

	body_str = string(body)

	return
}


func RquestFuke() (body_str string)  {
	url := "https://api.yuantutech.com/user-web/restapi/common/reservation/listScheduleinfoByDate?corpId=261&deptCode=0122&regMode=1&unionId=29&invokerChannel=H5&invokerDeviceType=weixin&invokerAppVersion=2.1.0&t=57500&callback=jsonp1555682519344"
	//resp, err := http.Get(url)
	//if err != nil {
	//	// handle error
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//
	//fmt.Println(string(body))
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("返回错误")
	}

	//还可以这样设置header， 另一种设置
	//req.Headers = map[string]string{
	//	"Accept-Encoding": "",
	//}

	req.Header.Add("Host", "api.yuantutech.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 7.1.1; MI MAX 2 Build/NMF26F; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/044611 Mobile Safari/537.36 MMWEBID/9053 MicroMessenger/7.0.3.1400(0x2700033C) Process/tools NetType/WIFI Language/zh_CN`)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Referer", "https://s.yuantutech.com/yuantu/h5-cli/2.1.0/appointment-select.html?deptCode=0080&corpId=261&deptName=%E4%BA%A7%E7%A7%91%E9%97%A8%E8%AF%8A&parentDeptCode=&regMode=1&regType=1&unionId=29&target=_blank")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br") //需要注释掉， 否则乱码
	req.Header.Add("Accept-Language", "zh-CN,en-US;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("返回错误")
	}
	body, err := ioutil.ReadAll(resp.Body)

	//body_str := string(body)
	//fmt.Println(string(body))

	body_str = string(body)

	return
}
