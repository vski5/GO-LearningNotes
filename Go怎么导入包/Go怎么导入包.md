### Go导入包

现在的情况是，我们在`moduledemo/main.go`中调用了`mypackage`这个包。

```bash
moduledemo
├── go.mod
├── main.go
└── mypackage
    └── mypackage.go
```

假设我们现在有`moduledemo`和`mypackage`两个包，其中`moduledemo`包中会导入`mypackage`包并使用它的`New`方法。

`mypackage/mypackage.go`内容如下：

```go
package mypackage  //调用mypackage包

import "fmt"

func New(){        //使用mypackage包里的New方法，注意此处N为大写。
	fmt.Println("mypackage.New")
}

```



先要建好.mod文件，见go的打包那个笔记。

目录结构在最上面

```go
package main

import (
	"fmt"
	"moduledemo/mypackage"  // 导入同一项目下的mypackage包，用的是相对位置，一个放项目的大文件夹里的一个包所在的文件夹。     此处还用了高并发特性（大概），在括号后面用两个引号来调用两个不同的包。
)
func main() {        //此处用了main函数，还使花括号下的两行在函数下的作用下执行。
	mypackage.New()    //调用了mypackage里的new功能
	fmt.Println("main") //调用了fmt函数里的println函数（用来打印输出结果的函数），打印main
}
```