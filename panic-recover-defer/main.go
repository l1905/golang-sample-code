package main

import "fmt"

/**
主要学习go panic-recover-defer异常

panic 和其他语言比如php的exception比较相像， 类似 throw exception
执行的主要流程是

进入函数

遇到panic,
1. 中止函数进行，
2. 调用函数defer

3. 返回调用者函数

进入调用者函数, 假装触发panic

1. 中止函数进行，
2. 调用函数defer
3. 继续返回上次调用者函数

...递归一次进行以上函数

当defer中遇到
recover, 则恢复停止panic流程， 上层函数继续执行正常逻辑， recover类似其他语言中的catch

如果函数存在多个defer: 原则上执行顺序是 "先"进"后"出， 参考链接： https://tour.golang.org/flowcontrol/13


参考链接：
https://blog.golang.org/defer-panic-and-recover
http://semicircle.github.io/blog/2013/09/24/letpanicfly/
https://studygolang.com/articles/14344?fr=sidebar

*/

//func main() {
//	f()
//	fmt.Println("Returned normally from f.")
//}
//
//func f() {
//	defer func() {
//		if r := recover(); r != nil {
//			fmt.Println("Recovered in f", r)
//		}
//	}()
//	fmt.Println("Calling g.")
//	g(0)
//	fmt.Println("Returned normally from g.")
//}
//
//func g(i int) {
//	if i > 3 {
//		fmt.Println("Panicking!")
//		panic(fmt.Sprintf("%v", i))
//	}
//	defer fmt.Println("Defer in g", i)
//	fmt.Println("Printing in g", i)
//	g(i + 1)
//}

func Process2() {
	CreatePanic()

	defer deferFunc2()
}

func Process() {
	Process2()
	defer deferFunc()
}

func deferFunc() {
	if r := recover(); r != nil {
		fmt.Println("recover========")
	}
	fmt.Println("deferFunc=========")
}

func deferFunc2() {
	fmt.Println("deferFun2============")
}


func CreatePanic() {
	fmt.Println("createPanic")
	defer deferPanic()
}

func deferPanic() {
	if r := recover(); r != nil {
		a := r.(string)
		fmt.Println(a)
		fmt.Println("recover========")

	}
}

func main() {
	Process()
}
