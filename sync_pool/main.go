package main

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"sync"
)

const MaxFrameSize = 5000


// put 两次，
// get两次 ，都能获取到值

var m sync.WaitGroup

func main() {
	data := (1 << 32) / 4
	fmt.Print(data)
	runtime.GOMAXPROCS(1)
	//getBuf()

	for i := 0; i < 500; i++ {
		m.Add(1)
		go func() {
			b := []byte("hello")
			z:=bytes.NewBuffer(b)
			putBuf(z)
			m.Done()
		}()
	}
	m.Wait()

	for i := 0; i < 500; i++ {
		m.Add(1)
		go func() {
			z := getBuf()
			fmt.Println(z.String())
			m.Done()
		}()
	}

	m.Wait()

	////runtime.GOMAXPROCS(3)
	//for i := 0; i < 10; i++ {
	//	// 多个协程想从pool中拿到对象
	//	go func() {
	//		c := getBuf()
	//		putBuf(c)
	//		log.Println("put done")
	//	}()
	//}
	//
	//time.Sleep(3 * time.Second)
}

var bufPool = sync.Pool{
	New: func() interface{} {
		log.Println("alloc=========")
		return bytes.NewBuffer(make([]byte, 0, MaxFrameSize))
	},
}

var bufPoolChan = make(chan bool, 1)

func getBuf() *bytes.Buffer {
	fmt.Println("getBuf...")
	b := bufPool.Get().(*bytes.Buffer)
	return b
}

func putBuf(b *bytes.Buffer) {
	fmt.Println("putBuf...")
	bufPool.Put(b)
}