# go-52-108-垃圾回收和 SetFinalizer

GC(垃圾收集器)：会收集程序中不再使用的变量和结构占用的内存。
可通过`runtime`包访问GC进程。

调用`runtime.GC()`函数会触发显示的GC，但这只在某些罕见的场景下才有用，
比如：当内存资源不足时调用 `runtime.GC()`，它会在此函数执行的点上立即释放一大片内存，此时程序可能会有短时的性能下降（因为 `GC` 进程在执行）。


# 如果想知道当前的内存状态，可以使用：
### func [ReadMemStats](https://github.com/golang/go/blob/master/src/runtime/mem.go?name=release#72 "View Source")
```go
func ReadMemStats(m *MemStats)
```
方法`ReadMemStats()`用`*MemStats`将内存申请和分配的统计信息填写进m
#### MemStats
`*MemStats` 是结构体MemStats的内容。
MemStats记录内存申请和分配的统计信息。
```go
type MemStats struct {
    // 一般统计
    Alloc      uint64 // 已申请且仍在使用的字节数
    TotalAlloc uint64 // 已申请的总字节数（已释放的部分也算在内）
    Sys        uint64 // 从系统中获取的字节数（下面XxxSys之和）
    Lookups    uint64 // 指针查找的次数
    Mallocs    uint64 // 申请内存的次数
    Frees      uint64 // 释放内存的次数
    // 主分配堆统计
    HeapAlloc    uint64 // 已申请且仍在使用的字节数
    HeapSys      uint64 // 从系统中获取的字节数
    HeapIdle     uint64 // 闲置span中的字节数
    HeapInuse    uint64 // 非闲置span中的字节数
    HeapReleased uint64 // 释放到系统的字节数
    HeapObjects  uint64 // 已分配对象的总个数
    // L低层次、大小固定的结构体分配器统计，Inuse为正在使用的字节数，Sys为从系统获取的字节数
    StackInuse  uint64 // 引导程序的堆栈
    StackSys    uint64
    MSpanInuse  uint64 // mspan结构体
    MSpanSys    uint64
    MCacheInuse uint64 // mcache结构体
    MCacheSys   uint64
    BuckHashSys uint64 // profile桶散列表
    GCSys       uint64 // GC元数据
    OtherSys    uint64 // 其他系统申请
    // 垃圾收集器统计
    NextGC       uint64 // 会在HeapAlloc字段到达该值（字节数）时运行下次GC
    LastGC       uint64 // 上次运行的绝对时间（纳秒）
    PauseTotalNs uint64
    PauseNs      [256]uint64 // 近期GC暂停时间的循环缓冲，最近一次在[(NumGC+255)%256]
    NumGC        uint32
    EnableGC     bool
    DebugGC      bool
    // 每次申请的字节数的统计，61是C代码中的尺寸分级数
    BySize [61]struct {
        Size    uint32
        Mallocs uint64
        Frees   uint64
    }
}
```

### 具体操作：
```go

/*
func ReadMemStats(m *MemStats)
ReadMemStats将MemStats(内存申请和分配的统计信息) 填写进m。
*/
var m runtime.MemStats  //将MemStats(内存申请和分配的统计信息) 填写进m
runtime.ReadMemStats(&m)  //与上面的func ReadMemStats(m *MemStats)作用相同
fmt.Printf("%d b\n", m.Alloc ) //Alloc是 已申请且仍在使用的字节数（就是内存大小单位是b）

```

MemStats记录内存申请和分配的统计信息。

ReadMemStats将内存申请和分配的统计信息填写进m。

上面的程序会给出已分配内存的总量，单位是b。

## func [SetFinalizer](https://github.com/golang/go/blob/master/src/runtime/extern.go?name=release#177 "View Source")
#### 函数用途：
如果需要在一个对象 obj 被从内存移除前执行一些特殊操作，比如写到日志文件中，就会用到func SetFinalizer()
#### 细则
```go
func SetFinalizer(x, f interface{})
```

SetFinalizer将x的终止器设置为f。

当垃圾收集器发现一个不能接触的（即引用计数为零，程序中不能再直接或间接访问该对象）具有终止器的块时，它会清理该关联（对象到终止器）并在独立go程调用f(x)。这使x再次可以接触，但没有了绑定的终止器。如果SetFinalizer没有被再次调用，下一次垃圾收集器将视x为不可接触的，并释放x。

SetFinalizer(x, nil)会清理任何绑定到x的终止器。

参数x必须是一个指向通过new申请的对象的指针，或者通过对复合字面值取址得到的指针。参数f必须是一个函数，它接受单个可以直接用x类型值赋值的参数，也可以有任意个被忽略的返回值。如果这两条任一条不被满足，本函数就会中断程序。

终止器会按依赖顺序执行：如果A指向B，两者都有终止器，且它们无法从其它方面接触，只有A的终止器执行；A被释放后，B的终止器就可以执行。如果一个循环结构包含一个具有终止器的块，该循环不能保证会被当垃圾收集，终止器也不能保证会执行；因为没有尊重依赖关系的顺序。

x的终止器会在x变为不可接触之后的任意时间被调度执行。不保证终止器会在程序退出前执行，因此一般终止器只用于在长期运行的程序中释放关联到某对象的非内存资源。例如，当一个程序丢弃一个os.File对象时没有调用其Close方法，该os.File对象可以使用终止器去关闭对应的操作系统文件描述符。但依靠终止器去刷新内存中的I/O缓冲如bufio.Writer是错误的，因为缓冲不会在程序退出时被刷新。

如果*x的大小为0字节，不保证终止器会执行。

一个程序会有单独一个go程顺序执行所有的终止器。如果一个终止器必须运行较长时间，它应该在内部另开go程执行该任务。

### 如果需要在一个对象 obj 被从内存移除前执行一些特殊操作，比如写到日志文件中，可以通过如下方式调用函数来实现：

```go
runtime.SetFinalizer(obj, func(obj *typeObj))
```

`func(obj *typeObj)` 需要一个 `typeObj` 类型的指针参数 `obj`，特殊操作会在它上面执行。`func` 也可以是一个匿名函数。

在对象被 GC 进程选中并从内存中移除以前，`SetFinalizer` 都不会执行，即使程序正常结束或者发生错误。