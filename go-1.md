# go-1

函数是基本的代码块，是执行一个任务时构成代码执行的逻辑结构。一个函数是在输入源基础上，通过执行一系列的算法，生成预期的输出

##### 函数的基本组成为：

1. 关键字func 

2. 函数名

3. 参数列表

4. 返回值

5. 函数体

6. 返回语句

   ### Hello World

   现在我们来创建第一个Go项目——`hello`。在我们桌面创建一个`hello`目录。

   #### go mod init

   使用go module模式新建项目时，我们需要通过`go mod init 项目名`命令对项目进行初始化，该命令会在项目根目录下生成`go.mod`文件。例如，我们使用`hello`作为我们第一个Go项目的名称，执行如下命令。

   ```bash
   go mod init hello
   ```

   #### 编写代码

   接下来在该目录中创建一个`main.go`文件：

```go
package main  
// 告诉语言，我调用了一个名为 main 的包，这个main包意味着当前是一个可执行程序。
//调用包的语言package.用来来指明调用的所在的包
//一个可执行程序有且仅有一个 main 包

import "fmt" 
// 导入内置 fmt 包
//用 import 来引入包时,默认是在引入上面package引用的文件夹里的GO包， fmt包含有格式化I/O函数

func main(){  
    // main函数，是程序执行的入口。function意思是功能，func也是一种函数，这一段意思是使用上面package提到的main包里的main方法，这段的main是上面提到的包里的一个方法或者叫做工具。关键字func .
    
	fmt.Println("Hello World!")  
    // 在终端打印 Hello World!。是调用fmt包里面的功能println，也就是打印。
}
```

### 编译

`go build`命令表示将源代码编译成可执行文件。

在hello目录下执行：

```bash
go build
```

或者在其他目录执行以下命令：

```bash
go build hello
```

go编译器会去 `GOPATH`的src目录下查找你要编译的`hello`项目

编译得到的可执行文件会保存在执行编译命令的当前目录下，如果是windows平台会在当前目录下找到`hello.exe`可执行文件。



# 还没来得及看





### Windows下VSCode切换cmd.exe作为默认终端

如果你打开VS Code的终端界面出现如下图场景（注意观察红框圈中部分），那么你的`VS Code`此时正使用`powershell`作为默认终端：![vscode shell配置1](https://www.liwenzhou.com/images/Go/install_go_dev/vscode_shell1.png)十分推荐你按照下面的步骤，选择`cmd.exe`作为默认的终端工具：![vscode shell配置2](https://www.liwenzhou.com/images/Go/install_go_dev/vscode_shell2.png)此时，VS Code正上方中间位置会弹出如下界面，参照下图挪动鼠标使光标选中后缀为`cmd.exe`的那一个，然后点击鼠标左键。

最后**重启VS Code中已经打开的终端**或者**直接重启VS Code**就可以了。![vscode shell配置3](https://www.liwenzhou.com/images/Go/install_go_dev/vscode_shell3.png)如果没有出现下拉三角，也没有关系，按下`Ctrl+Shift+P`，VS Code正上方会出现一个框，你按照下图输入`shell`，然后点击指定选项即可出现上面的界面了。![vscode shell配置4](https://www.liwenzhou.com/images/Go/install_go_dev/vscode_shell4.png)

### go run

`go run main.go`也可以执行程序，该命令本质上也是先编译再执行。

### go install

`go install`表示安装的意思，它先编译源代码得到可执行文件，然后将可执行文件移动到`GOPATH`的bin目录下。因为我们的环境变量中配置了`GOPATH`下的bin目录，所以我们就可以在任意地方直接执行可执行文件了。

### 跨平台编译

默认我们`go build`的可执行文件都是当前操作系统可执行的文件，如果我想在windows下编译一个linux下可执行文件，那需要怎么做呢？

只需要指定目标操作系统的平台和处理器架构即可，例如Windows平台cmd下按如下方式指定环境变量。

```bash
SET CGO_ENABLED=0  // 禁用CGO
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64
```

注意：如果你使用的是PowerShell终端，那么设置环境变量的语法为`$ENV:CGO_ENABLED=0`。

然后再执行`go build`命令，得到的就是能够在Linux平台运行的可执行文件了。

Mac 下编译 Linux 和 Windows平台 64位 可执行程序：

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

Linux 下编译 Mac 和 Windows 平台64位可执行程序：

```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

Windows下编译Mac平台64位可执行程序：

```bash
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build
```

现在，开启你的Go语言学习之旅吧。人生苦短，let’s Go.