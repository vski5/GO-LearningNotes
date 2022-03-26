# go-36-083-084-for-range 的配套用法-map 类型的切片

用for-range历遍map()

```go
for key,value := range map1 {

}
```

因为map()是无序的，所以每次历遍出来的顺序都不一样。

