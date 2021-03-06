# go-32-074-075-切片重组（reslice）-切片的复制与追加

# 切片重组

切片是数组的view，所以只会小于或者等于数组。

```go
slice1 := make([]type, start_length, capacity)
```

将切片扩展 1 位可以这么做：

**用`=`重赋值**

```go
sl = sl[0:len(sl)+1]
```

切片可以反复扩展直到占据整个相关数组,也就是到cap的极限。

这就是**切片重组**。



# 切片的复制与追加

假设 s 是一个字符串（本质上类似一个字节数组，但不可更改元素），

比如：

```go
s := make([]int,5,6)
```



1. 可以直接通过 `c := []byte(s)` 来获取单个字节的切片 c。

2. 也可以通过copy函数达到，

   ```go
   copy (dst []byte,src string)
   ```

   dst和src都是代指一个slice的函数名。

   

   golang的`builtin`包提供的`copy`方法可以用来复制切片。该方法返回成功复制元素的个数，签名如下：

   ```swift
   func copy(dst, src []Type) int
   ```

   被复制的元素个数是dst和src中短的那个。同时注意一旦复制，对dst的任何修改都不会影响到src，反之亦然。

3. 也可以通过for-range历遍实现，

   ```go
   package main
   
   import "fmt"
   
   func main() {
       s := "\u00ff\u754c"
       for i, c := range s {
           fmt.Printf("%d:%c ", i, c)
       }
   }
   ```

## 获取字符串的一部分

使用 `substr := str[start:end]` 可以从字符串 str 获取到从索引 start 开始到 `end-1` 位置的子字符串。

`str[start:]` 则表示获取从 start 开始到 `len(str)-1` 位置的子字符串。

而 `str[:end]` 表示获取从 0 开始到 `end-1` 的子字符串。

## 字符串和切片的内存结构

在内存中，一个字符串实际上是一个双字结构。

字符串 是一个指向实际数据的**指针**和记录字符串长度的整数。

因为指针对用户来说是完全不可见，所以，只看作为后半部分，也就是认为字符串是一个值类型，也就是一个字符数组。



字符串 `string s = "hello"` 和子字符串 `t = s[2:3]` 在内存中的结构可以用下图表示：

![](C:\Users\李尤的光影精灵.LAPTOP-85F2O9KH\Desktop\GOnotes\the-way-to-go_ZH_CN-master\eBook\images\7.6_fig7.4.png)



## 修改字符串中的某个字符

字符串不可变，只能曲线救国。

`s := "hello"`就是字符串。

`c := []byte(s)`字符串转化为字节数组。

`[]byte()`就是字节数组。

byte是字节，正常的数组是`[]type()`



方法： 

1. 先将字符串转化为字节数组，
2. 而后再修改数组中的元素值，
3. 最后将数组转换为字符串格式。

例子：

```GO
s := "hello"
c := []byte(s) //字符串转化为字节数组
c[0] = 'c' //修改字节数组中的第0个（首个）
s2 := string(c) // s2 == "cello" 最后将数组转换为字符串格式。
```

默写修改字节数组中的第0个（首个）：

```GO
c[0] = 'p'
```

默写字符串转化为字节数组:

```GO
s := "hello"
c := []byte(s)
```



## 字节数组对比函数

就是按位比较整数。

用`switch-case`语句解决

## 搜索 及 排序  切片和数组

`sort`包用于搜索和排序。

这就是如何使用 `sort` 包的方法，我们会在第 11.6 节对它的细节进行深入，并实现一个属于我们自己的版本。

## append 函数常见操作

**append：**

- **v.**增补
- **网络**附加；追加；添加

作用：

1. 将切片b的元素追加到切片a中：`s := append(a,b)`

   据此特性可以做到删除位于索引i的元素，因为切片右开左闭：`s := append(s[:i],s[i+1:])`

   append()拼接括号内的一切数组，所以可以用来插入拼接修改各种数组

2. 复制切片 a 的元素到新的切片 b 上：

   ```go
   b = make([]T, len(a))
   copy(b, a)
   ```




**一个复习**

make()的使用方法是：

```go
func make([]T, len, cap)
/*
func make([]类型，长度，最大长度)
*/
```

new()的使用方法是：

```go
new([100]int)[0:50]
/*
func new([最大长度]类型)[起始位置:终结位置]
*/
```



#  切片和垃圾回收

我在下一节会详细写。



