# go-46-102-工厂方法（method） 与 结构体实例

类似于go-45中的

```go
func newPerson(name1 string) *person {
	p := person{name: name1}
	p.age = 42
	return &p 
}
```

就是制造一个用于赋值到结构体的工厂。

## 结构体工厂

GO不支持构建子方法，只能用“构建子工厂”的方法解决。

一般为类型定义一个工厂。

惯例：工厂名字以new（或NEW类似的）开头。

假设定义一个file结构体类型：

```go
type file struct{
    fd int //文件描述符
    name string //文件名
}
```

为这个file结构体类型定义一个工厂方法，

这个方法会返回一个指向结构体实例的指针：

```GO
func newfile (fd int , name string) *file{
    if fd<0 {
        return nil
    }
    return &file{fd , name}
}
```

指针指向内存所在地，可以直接从内存中读数据，这相当于直接从门牌号中查水表。

**这个方法的使用方式是：**

```GO
something := newfile{10,"./test.txt"}
```

用了之后就会返回`&file{10,"./test.txt"}`。

请注意，在file是结构体类型的情况下：`new(file)`和`&file{}`是等价的。

如何查结构体类型ST的一个实例占用的内存：`size1 := unsafe.Sizeof(ST{})`。

## 在其他包里使用自己创造的工厂方法，又称强制使用工厂方法。

要写清楚调用的是哪儿的方法，专业术语来说就是：选择器要写得清楚明白。

params是参数的意思，此处代指输入值；matrix是模型、矩阵的意思。

在matrix文件中创建一个NewMatrix方法：

```go
type matrix struct{
    
}
func NewMatrix(params) *matrix{
    m := new(matrix) //初始化m
    return m
}
```

在其他包里使用：

```GO
package main 
import ("matrix")
RightWay := martrix.NewMatrix(params)
```



# map和struct     VS     new()和make()

一个复习：

```GO
m := make([]type,len,cap)
n := new([cap]type)[10:20]
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



make()有三个类型：slice / maps / channels







































