package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func main()  {
	ExceOnce()
}

func padRightSide(str string, item string, count int) string {
	return str + strings.Repeat(item, count)
}

func load() {
	a := "1"

	fmt.Printf("%04s", a)
	fmt.Printf("Ending or pad RIGHT with zero :  |%v|\n", padRightSide("108", "0", 3))
}


// 更新waitMethod
func waitMethod() {
	var wg sync.WaitGroup

	wg.Add(-1)

	go func() {
		fmt.Println("Goroutine 1")
		//wg.Add(-1)
		wg.Done()
	}()

	go func() {
		fmt.Println("Goroutine 2")
		time.Sleep(1 * time.Second)
		//wg.Add(-1)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("进程结束")
}


var once sync.Once
var singleton *Singleton

type Singleton struct {
	frist string
}

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{frist: "初始化"}
	})

	return singleton
}

// 单例模式
func ExceOnce()  {
	var s *Singleton
	s = GetInstance()
	s = GetInstance()

	fmt.Println(s.frist)
}