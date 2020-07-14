package guahao

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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
	client := &http.Client{
		Timeout: time.Second * 3,
	}

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
	client := &http.Client{
		Timeout: time.Second * 3,
	}

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

//儿童早期发展中心
func RquestFazhan() (body_str string)  {
	url := "https://api.yuantutech.com/user-web/restapi/common/reservation/listScheduleinfoByDate?corpId=261&deptCode=0009&regMode=1&unionId=29&invokerChannel=H5&invokerDeviceType=weixin&invokerAppVersion=2.12.4&t=72961&callback=jsonp1583748400816"
	client := &http.Client{
		Timeout: time.Second * 3,
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("返回错误")
	}

	req.Header.Add("Host", "api.yuantutech.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 7.1.1; MI MAX 2 Build/NMF26F; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/044611 Mobile Safari/537.36 MMWEBID/9053 MicroMessenger/7.0.3.1400(0x2700033C) Process/tools NetType/WIFI Language/zh_CN`)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Referer", "https://s.yuantutech.com/yuantu/h5-cli/2.12.4/appointment-select.html?deptCode=0009&corpId=261&deptName=%E5%84%BF%E4%BF%9D-%E6%97%A9%E6%9C%9F%E5%8F%91%E5%B1%95%E9%97%A8%E8%AF%8A&parentDeptCode=&regMode=1&regType=1&unionId=29&target=_blank&spm=100.sections.dept.0009")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br") //需要注释掉， 否则乱码
	req.Header.Add("Accept-Language", "zh-CN,en-US;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("返回错误")
	}
	fmt.Println(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)

	//body_str := string(body)
	//fmt.Println(string(body))

	body_str = string(body)

	return
}

//儿童生长发育门诊
func RquestFayu() (body_str string)  {
	url := "https://api.yuantutech.com/user-web/restapi/common/reservation/listScheduleinfoByDate?corpId=261&deptCode=0390&regMode=1&unionId=29&invokerChannel=H5&invokerDeviceType=weixin&invokerAppVersion=2.16.6&t=47296&callback=jsonp1586267700968"
	client := &http.Client{
		Timeout: time.Second * 3,
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("返回错误")
	}

	req.Header.Add("Host", "api.yuantutech.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 7.1.1; MI MAX 2 Build/NMF26F; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/044611 Mobile Safari/537.36 MMWEBID/9053 MicroMessenger/7.0.3.1400(0x2700033C) Process/tools NetType/WIFI Language/zh_CN`)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Referer", "https://s.yuantutech.com/yuantu/h5-cli/2.12.4/appointment-select.html?deptCode=0009&corpId=261&deptName=%E5%84%BF%E4%BF%9D-%E6%97%A9%E6%9C%9F%E5%8F%91%E5%B1%95%E9%97%A8%E8%AF%8A&parentDeptCode=&regMode=1&regType=1&unionId=29&target=_blank&spm=100.sections.dept.0009")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br") //需要注释掉， 否则乱码
	req.Header.Add("Accept-Language", "zh-CN,en-US;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("返回错误")
	}
	fmt.Println(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)

	//body_str := string(body)
	fmt.Println(string(body))

	body_str = string(body)

	return
}


// 黄岛妇产科
func RquestHuangDaoFuchan() (body_str string)  {
	url := "https://api.yuantutech.com/user-web/restapi/common/reservation/listScheduleinfoByDate?corpId=29003&deptCode=558&regMode=1&unionId=29&invokerChannel=H5&invokerDeviceType=weixin&invokerAppVersion=2.17.13&t=63694&callback=jsonp1591757673918"
	client := &http.Client{
		Timeout: time.Second * 3,
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("返回错误")
	}

	req.Header.Add("Host", "api.yuantutech.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", `Mozilla/5.0 (iPhone; CPU iPhone OS 13_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/7.0.12(0x17000c30) NetType/WIFI Language/zh_CN`)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Referer", "https://s.yuantutech.com/yuantu/h5-cli/2.17.13/appointment-select.html?deptCode=558&corpId=29003&deptName=%E8%A5%BF%E6%B5%B7%E5%B2%B8%E4%BA%A7%E7%A7%91%E9%97%A8%E8%AF%8A&parentDeptCode=&regMode=1&regType=1&unionId=29&target=_blank&spm=100.sections.dept.558")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br") //需要注释掉， 否则乱码
	req.Header.Add("Accept-Language", "zh-CN,en-US;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("返回错误")
	}
	fmt.Println(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)

	//body_str := string(body)
	fmt.Println(string(body))

	body_str = string(body)

	return
}