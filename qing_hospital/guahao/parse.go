package guahao

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

type Resp struct {
	Success bool `json:"success"`
	ResultCode string `json:"resultCode"`
	Msg string `json:"msg"`
	Data []RespData `json:"data"`
	StartTime int64 `json:"startTime"`
	TimeConsum int `json:"timeConsum"`
}

type RespData struct {
	CorpName string `json:"corpName"` //tag标记，可以指定json对象的字段名称
	Date string `json:"date"` //如果Nickname字段为空，则忽略转化, 中间","不能有空格
	DeptName      int `json:"deptName"` //这里重命名为其他名字
	PanYuMode string `json:"panYuMode"` //序列化时， 不输出, 和字段小写不输出类似
	ScheduleTypeVOList []ScheduleTypeVOList `json:scheduleTypeVOList`
}

type ScheduleTypeVOList struct {
	Name string `json:"name"`
	DoctorVOList []doctorVOList `json:"doctorVOList"`

}

type doctorVOList struct {
	Name string `json:"name"`
	DeptName string `json:"deptName"`
	DoctCode string `json:"doctCode"`
	DeptCode string `json:"deptCode"`
	DoctLogo string `json:"doctLogo"`
	MedDate string `json:"medDate"`
	MedStatusDesc string `json:"medStatusDesc"`
	MedAmNum int `json:"medAmNum"`
	MedPmNum int `json:"medPmNum"`
	RegAmount string `json:"regAmount"`
	DoctTech string `json:"doctTech"`
	Sex string `json:"sex"`
	RegTypeName string `json:"regTypeName"`
	ScheduleId string `json:"scheduleId"`
	RegType int `json:"regType"`
	AmScheduleId string `json:"amScheduleId"`
	PmScheduleId string `json:"pmScheduleId"`
}

func Regex(text string, date_zone string, time string) (print_text string)  {
	formatRegex := regexp.MustCompile(`jsonp\d+\((.*?)\);`)
	//类型推算使用
	params := formatRegex.FindStringSubmatch(text)
	if params != nil {
		real_text := []byte(params[1])

		var resp Resp
		err := json.Unmarshal(real_text, &resp)

		if err != nil {
			//
		}
		//fmt.Printf("%+v\n", resp) //%+v格式化输出 golang对象
		//fmt.Println(resp.Data[0].CorpName)
		data_list := resp.Data


		//date_zone_list := []{}
		for _, day_detail := range data_list {
			day_list := day_detail.ScheduleTypeVOList

			for _, day_kind := range day_list {
				for _, row := range day_kind.DoctorVOList {
					//fmt.Printf("%+v\n", row.MedAmNum)
					//fmt.Printf("%+v\n", row.MedPmNum)
					//1. 时间区间过滤, 正则匹配
					date_zone_regex := regexp.MustCompile(day_detail.Date)

					date_check := false
					if len(date_zone) > 0 && date_zone_regex.MatchString(date_zone)  {
						date_check = true
					}
					if len(date_zone) == 0 {
						date_check = true
					}
					if date_check {
						name := row.Name
						sex := row.Sex

						if name != "计划生育" {
							all_day_num_text := ""

							//日期过滤
							if (len(time) == 0 || time == "AM") &&  row.MedAmNum > 0 {
								all_day_num_text += " 上午还剩"+ strconv.Itoa(row.MedAmNum)
							}
							if (len(time) == 0 || time == "PM") && row.MedPmNum > 0 {
								all_day_num_text += " 下午还剩" + strconv.Itoa(row.MedPmNum)
							}
							if len(all_day_num_text) > 0 {
								print_text +=  fmt.Sprintf("%s:%s:%s:%s:%s:%d:%s\n", row.DeptName, day_detail.Date, name, sex, row.DoctTech, row.RegType, all_day_num_text)
							}
						}
					}



				}
			}
		}
		//fmt.Println(print_text)
	}

	return


}
