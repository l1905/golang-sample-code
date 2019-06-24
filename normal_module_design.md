通用模块设计， 比如 日志， stat 统计模块

```

var handler

func newInit() {
    handler = &{}
}

func Error() {
    handler.errorMethod()
}

func Notice() {
    handler.errorMethod()
}

调用方法

conf.newInit()

conf.Error('xxx')


第二种

func newInit() {
    return &conf{}
}

(confHandler *conf ) func Error() {
    confHandler.errorMethod()
}

(confHandler *conf ) func Notice() {
    confHandler.errorMethod()
}

使用方式
confHandler = conf.newInit()
confHandler.Error('xxxxx')

```

原理都是单例模式， 但是下面这个，依赖需要传递 指针
