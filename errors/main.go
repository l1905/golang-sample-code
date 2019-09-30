package main

import (
	"errors"
	"fmt"
)



/*
// 定义接口
type error interface {
	Error() string
}
*/

// golang version 1.13+
// 参考文章： https://mp.weixin.qq.com/s/oqwwec_sTjfowhXYi5kiQg

// 错误类型
type MyError struct {
	name string
	value string
}

// 实现
func (my *MyError) Error() (string) {
	return fmt.Sprintf("name:%s,12345 value:%s", my.name, my.value)
}

func main() {

	// 抽象出方法
	// 实现Error()方法的Error结构体对象
	var err = &MyError{name:"litongxu", value:"100"}

	// 外面包上新的string Error
	newErr := fmt.Errorf("新的error:%w", err)

	// 获取里面一层的err信息
	inErr := errors.Unwrap(newErr)
	fmt.Println(inErr)

	// 判断newErr 是 err 类型
	fmt.Println(errors.Is(newErr, err))

	// AS 判断是否是这个类型， 并且, wrap里面存在，则将进行赋值操作
	var descError *MyError
	if errors.As(err, &descError) {
		fmt.Println("AS 操作成功")
	} else {
		fmt.Println("AS 操作失败")
	}

	// 类似interface 的类型推断方法.(xxxTYPE)
	fmt.Println(descError)

}
