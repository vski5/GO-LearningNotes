# go-36-083-084-for-range 的配套用法-map 类型的切片

# for-range 的配套用法

用for-range历遍map()

```go
for key,value := range map1 {

}
```

因为map()是无序的，所以每次历遍出来的顺序都不一样。

```go
package main

import "fmt"

func main() {
	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cris"}

	//range 在 map 中迭代键值对。
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range 也可以只遍历 map 的键。
	for k := range kvs {
		fmt.Println("key:", k)
	}
	//range也能只历遍map的值
	for _, value := range kvs {
		fmt.Println("value:", value)
	}

}

```

输出：

```
b -> banana
c -> cris
a -> apple
key: a
key: b
key: c
value: apple
value: banana
value: cris
```

# map 类型的切片

获得一个map类型的切片需要两次make()函数。

```go
package main
import "fmt"

func main() {
	// Version A:
	items := make([]map[int]int, 5)
	for i:= range items {
		items[i] = make(map[int]int, 1)//通过索引使用切片的 map 元素
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	// Version B: NOT GOOD!
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1) // item is only a copy of the slice element.
		item[1] = 2 // This 'item' will be lost on the next iteration.
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)
}
```

