## atomic变量相关


原子变量相关， 避免使用锁， 性能更好。

主要再以下三个场景使用

1. 设置原子变量 update
2. 增量原子变量 incr,decr
3. 比较大小，并且赋值
5. 加载获取变量值 load


loadPointer 和 storePointer最本质的区别

是希望用新地址，覆盖老地址，即新老地址不一样

如果继续使用