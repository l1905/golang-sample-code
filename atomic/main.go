package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

var intData int32
func main() {
	// 指针
	testPoiner()
}

var info unsafe.Pointer = nil
// 测试
type data struct {
	name string
	age int
}
func testPoiner()  {

	i := &data{
		name: "李同",
		age: 18,
	}

	atomic.StorePointer(&info, (unsafe.Pointer(i)))


	d := (*data)(atomic.LoadPointer(&info))

	fmt.Printf("%+v, %p \n", d, d)

	fmt.Printf("%p, \n", i)

	d.age = 1122434


	m := (*data)(atomic.LoadPointer(&info))

	fmt.Println(m)


}


func testInt() {
	intUpdate()
	intIncr()
	intSwap()
	fmt.Println(intLoad())
}

func intUpdate() {
	atomic.StoreInt32(&intData, 10)
}

func intIncr() {
	atomic.AddInt32(&intData, 1)
}

func intSwap() {
	atomic.CompareAndSwapInt32(&intData, 77, 88)
}

func intLoad() (int32) {
	return atomic.LoadInt32(&intData)
}
