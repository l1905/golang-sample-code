package main

import (
	"fmt"
	"time"
)

func main() {
	// 无缓存队列
	//uncacheChannel()

	// 单向队列测试
	/*singleChannel := make(chan string)
	go singleDirectSendChannel(singleChannel)
	//阻塞接收
	singleData := <-singleChannel
	fmt.Println(singleData)*/


	// 发布队列
	// channel已经被close， 继续发送消息会引发panic
	//sendPanicWhenClose()

	// chanel已经被重制为 nil, 继续发送消息，会引起阻塞
	//sendBlockWhenNil()

	// 当channel关闭，继续消费时， 会出现零值情况
	//consumeZeroValueWhenClose()

	// 从一个nil channel中接收数据会一直被block
	//consumeBlockWhenNil()

	// 通过range进行消费
	//consumeByRange()

	// select 阻塞实践， 即没有default, 会一直阻塞
	//selectBlock()

	// select 非阻塞实践, 存在default
	//selectUnBlock()

	// time实践 time.after 每xx时间，做某些事情， 类似 ticker
	//timeAfter()

	//timer 和stop搭配使用
	//timerWithStop()

	// timer 和reset 搭配使用
	//timerWithReset()

	// ticker test
	//tickerWithRange()

	time.Sleep(100*time.Second)
}

// 无缓存队列
func uncacheChannel() {
	chUnCache := make(chan float64)

	go func() {
		chUnCache<- 1;
	}()

	//阻塞接收协助程中的数据
	data := <-chUnCache

	fmt.Println(data)
}

// 定义参数是单向队列， 即在方法里面只能接收消息，限制使用
func singleDirectSendChannel(single chan <- string) {

	single <- "12345"
}

func sendPanicWhenClose() {
	chUnCache := make(chan float64, 100)

	go func() {
		chUnCache <- 1
		time.Sleep(1*time.Second)
	}()

	close(chUnCache)
}

func sendBlockWhenNil() {
	chUnCache := make(chan float64, 100)

	go func() {

		for i := 0; i< 100; i++ {
			chUnCache <- 1
			fmt.Println("1")
			time.Sleep(1*time.Second)
		}

	}()
	// 发送五次之后，阻塞发送进程
	time.Sleep(5*time.Second)
	chUnCache = nil
}

func consumeZeroValueWhenClose() {
	chUnCache := make(chan float64, 100)

	go func() {
		for i := 0; i< 100; i++ {
			data, ok := <-chUnCache

			// 当ok为false时， 可以判断channel已经关闭 todo 额外验证
			fmt.Println("===")
			time.Sleep(1*time.Second)

			fmt.Println(data)
			fmt.Println(ok)
		}
	}()
	chUnCache <- 99
	//chUnCache <- 0
	close(chUnCache)

/*	===
	99
	true
	===
	0
	false
	===
	0
	false
	===
	0
	false*/

}

func consumeBlockWhenNil()  {

	chUnCache := make(chan float64, 100)

	go func() {
		for i := 0; i< 100; i++ {
			// 执行五次之后， 会一直阻塞
			data, ok := <-chUnCache

			// 当ok为false时， 可以判断channel已经关闭 todo 额外验证
			fmt.Println("===")
			time.Sleep(1*time.Second)

			fmt.Println(data)
			fmt.Println(ok)
		}
	}()

	for i := 0; i< 100; i++ {
		chUnCache <- float64(i)
	}

	//chUnCache <- 0
	time.Sleep(5*time.Second)
	chUnCache = nil
}

func consumeByRange()  {
	go func() {
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		// 只有close关掉后， range才会结束， 否则会一直阻塞
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Finished")
}


// select阻塞实践
func selectBlock() {
	c := make(chan int, 100)

	go func() {
		time.Sleep(5 * time.Second)

		for i:=0; i < 10; i++ {
			c <- i
		}
	}()

	for {
		println("start----for")
		select {
		case i := <-c:
			fmt.Println(i)
		}
		println("start----end")
	}

	// fatal error: all goroutines are asleep - deadlock!
	// 最后会报这个错误， 因为协程退出，主进程的select, 永远不会等到消息, 解决办法是不让协程退出，或者使用default

}

// select阻塞实践
func selectUnBlock() {
	c := make(chan int, 100)

	go func() {
		time.Sleep(5 * time.Second)

		for i:=0; i < 10; i++ {
			c <- i
		}
	}()

	for {
		println("start----for")
		select {
		case i := <-c:
			fmt.Println(i)
		default:
			fmt.Println("default")
		}
		println("start----end")
	}

}

func timeAfter() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	//time.after 内部实现： NewTimer(d).C， 不会释放垃圾回收， 建议使用 NewTimer 并且stop操作
}

// 使用stop 停止timer
func timerWithStop()  {
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

func timerWithReset()  {
	t := time.NewTimer(time.Second * 2)
	defer t.Stop()
	for {
		<-t.C
		fmt.Println("timer running...")
		// 需要重置Reset 使 t 重新开始计时, 如果没有Reset 会导致 fatal error: all goroutines are asleep - deadlock!， 因为t.c
		// 永远获取不到数据
		// defer t.Stop()在这里并不会停止定时器。这是因为Stop会停止Timer，
		// 停止后，Timer不会再被发送，但是Stop不会关闭通道，防止读取通道发生错误。
		t.Reset(time.Second * 2)
	}
}

func tickerWithRange()  {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
}


// 使用场景

// for -range消费

