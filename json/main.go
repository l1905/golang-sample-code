package main

/**
golang学习，常用代码练习
json相关基本操作
参考链接
https://sanyuesha.com/2018/05/07/go-json/
https://eager.io/blog/go-and-json/
https://golang.org/pkg/encoding/json/


核心点：
数据结构序列化为json
json反序列化为数据结构

 */

import (
	"encoding/json"
	"fmt"
)

type Bank struct {
	Name string `json:"name"`
	Address string `json:"address"`
}

type User struct {
	UserName string `json:"username"` //tag标记，可以指定json对象的字段名称
	NickName string `json:"nickname,omitempty"` //如果Nickname字段为空，则忽略转化, 中间","不能有空格
	Age      int `json:"ageage"` //这里重命名为其他名字
	Birthday string `json:"-"` //序列化时， 不输出, 和字段小写不输出类似
	BankInfo Bank `json:"bank_info"` //嵌套json 序列
	Sex      string
	Email    string  `json:"email"` //嵌套json 序列
	Phone    string

	city     string //只能序列化exported filed, 即首字母大写， 该字段小写，私有，会被过滤掉
	Country	 string
}

/*结构体转json*/
func struct2Json() {
	fmt.Println("struct2Json======")
	user1 := &User{
		UserName: "user1",
		//NickName: "",
		//Age:      18, //这里如果不存在话， 会输出默认值
		Birthday: "2008/8/8",
		BankInfo: Bank{"青岛", "山东路"},
		Sex:      "男",
		Email:    "mahuateng@qq.com",
		Phone:    "110",
		city:     "青岛",
		Country:   " 中国",
	}
	//序列化
	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func json2Struct() {
	fmt.Println("json2Struct======")
	type App struct {
		Id string `json:"id"`
		Title string `json:"title"`
	}

	data := []byte(`
    {
        "id": "k34rAT4",
        "title": "My Awesome App"
    }
`)

	var app App
	//反序列化
	err := json.Unmarshal(data, &app)

	if err != nil {
		return
	}
	fmt.Printf("%+v\n", app) //%+v格式化输出 golang对象
}

//我们序列化之前不知道其数据格式，我们可以使用 interface{} 来存储我们的 decode 之后的数据
func json2Interface()  {
	fmt.Println("output json2Interface======")
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		return
	}
	m := f.(map[string]interface{}) //这里是类型推断，

	//遍历输出结果
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

func twice()  {
	// 一个对象可以连续两次作为unmarshal,的第二个参数， 新值只会更新 新字段，而不会完全覆盖
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	user1 := &User{}
	json.Unmarshal(b, user1)
	fmt.Printf("%+v \n", user1)
	c := []byte(`{"email":"85@qq.com"}`)
	json.Unmarshal(c, user1)
	fmt.Printf("%+v \n", user1)
}

func main() {
	twice()
	//struct2Json()
	//json2Struct()
	//json2Interface()
	fmt.Println("----")
}
