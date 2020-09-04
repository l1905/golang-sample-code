package main

import "fmt"

// 直接通过挂载的方法洗修改原始结构体，
type basic struct {
	UserName string
	City string
}

func (b *basic) Work() {
	// 覆盖结构体，生效
	*b = basic{
		UserName: "",
		City:     "",
	}
}

func (b *basic) NoWork() {
	// 覆盖结构体，生效
	b = &basic{
		UserName: "",
		City:     "",
	}
}

func main() {
	b := basic{
		UserName: "litong",
		City:     "qingdao",
	}

	b.Work()

	fmt.Println(b.UserName)
	fmt.Println(b.City)

}

