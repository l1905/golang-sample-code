package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("测试promot方法")
	prompt()

	fmt.Println("暂停后继续执行")

	fmt.Println("%s", int32Ptr(123))

	//先获取指针，然后再获取指针地址对应的具体值
	fmt.Println("%s", *int32Ptr(123))

}

//暂停等待输出, 可以暂停查看当前打印输出， 摘抄自k8s官网demo
func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}

//返回32位变量地址, 摘抄自k8s官网demo
func int32Ptr(i int32) *int32 { return &i }
