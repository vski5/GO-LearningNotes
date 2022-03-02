# go-17-select关键字用法

select 的意思是 选择、挑选。

select基本用法

```go
select {
case <- chan1:
// 如果chan1成功读到数据，则进行该case处理语句
case chan2 <- 1:
// 如果成功向chan2写入数据，则进行该case处理语句
default:
// 如果上面都没有成功，则进入default处理流程
```

这个结构就有点像switch结构。



`select`结构的规则如下：

- 在`select`语句中可以有任意数量的`case`语句。每个`case`语句后面都跟要比较的值和冒号。

- 每个 case 都必须是一个通信。

  因为select为通信而设计，每个case表达式中只能包含操作通道的表达式。

- 当发生通道操作时，将执行该`case`语句之后的语句。 在`case`语句中可不需要`break`语句。
- `select`结构后面的`default`语句可有可无，但它必须出现在`select`语句的结尾。
- `default`语句可用于在没有任何`case`语句为真时执行任务。在`default`语句不需要`break`语句。
- 就像是有个通向很多直流的河，所有case后面的语句不可运行时，也没有default语句，那么就会等到case后面的条件成立可运行为止。



# 未完待续

