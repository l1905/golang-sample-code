
## sync包主要方法学习


### `sync.wait`包使用

在 `sync.WaitGroup`（等待组）类型中，每个 `sync.WaitGroup` 值在内部维护着一个计数，此计数的初始默认值为零。


它提供了三个方法，Add()用来添加计数。Done()用来在操作结束时调用，使计数减一。Wait()用来等待所有的操作结束，即计数变为0，该函数会在计数不为0时等待，在计数为0时立即返回。


常用使用方式

```

var wg *sync.WaitGroup

wg.Add(1)
go func() {
    // logic

    wg.Done()
}

wg.Wait()

```

这样可以将异步函数调用，变成同步

需要额外注意的点

1. 计数器初始化值不能为负值
2.  WaitGroup对象不是一个引用类型  
   ```
     func main() {
         wg := sync.WaitGroup{}
         wg.Add(100)
         for i := 0; i < 100; i++ {
             go f(i, &wg)
         }
         wg.Wait()
     }
     
     // 一定要通过指针传值，不然进程会进入死锁状态
     func f(i int, wg *sync.WaitGroup) { 
         fmt.Println(i)
         wg.Done()
     }
   ```


### `sync.once`处理单例模式

在初始化全局变量时， 方便使用


