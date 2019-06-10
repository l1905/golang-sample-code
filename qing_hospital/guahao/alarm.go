package guahao

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const dingUrl = "https://oapi.dingtalk.com/robot/send?access_token=dea1327c54f40fb2d71d8f292a7ebf84d5ba46fa12cac6b2b638c5b91748a31b8" //调整为自己的token
const weixinUrl = "https://sc.ftqq.com/SCU110894T7d730de004034a8f8de6eca68f9ceb485993ab6b610fc.send"                                  //调整为自己的token

//调用app参数
type params struct {
	MsgType string      `json:"msgtype"`
	Text    params_text `json:"text"`
	At      params_at   `json:"at"`
}

type params_text struct {
	Content string `json:"content"`
}

type params_at struct {
	AtMobiles string `json:"atMobiles"`
}

func SendAlarm(message string) {
	if len(message) <= 0 {
		fmt.Println("没有数据")
		return
	}
	DingAlarm(message)
	WeixinAlarm(message)
}

func DingAlarm(message string) {
	if len(message) <= 0 {
		return
	}
	data := params{
		MsgType: "text",
		Text:    params_text{Content: message},
		At:      params_at{AtMobiles: "18765900000"}, //替换成自己的手机号
	}

	jsonStr, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", dingUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func WeixinAlarm(message string) {
	if len(message) <= 0 {
		return
	}
	t := time.Now().Local()
	s := t.Format("2006-01-02")

	message += ";当前时间:" + s
	message = strings.Replace(message, ":", " ", -1)
	message = strings.Replace(message, ",", " ", -1)
	message = strings.Replace(message, "\n", "=======", -1)

	client := &http.Client{}

	req, err := http.NewRequest("GET", weixinUrl, nil)

	if err != nil {
		fmt.Println("返回错误")
	}
	q := req.URL.Query()
	q.Add("text", "挂号报警"+s)
	message_len := 0
	if len(message) > 1000 {
		message_len = 1000
	} else {
		message_len = len(message)
	}
	q.Add("desp", message[:message_len])
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("返回错误")
	}
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(body)

}
