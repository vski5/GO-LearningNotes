# go-44-098-自定义包的目录结构、go install 和 go test

假设：有一个名为`uc`的包，其中含有一个 `UpperCase` 函数将字符串的所有字母转换为大写。

# 自定义包的目录结构

uc 代表通用包名, 

名字为粗体的代表目录，

斜体代表可执行文件

```
/home/user/goprograms —— 也就是 环境变量 GOPATH
│
├──ucmain.go	(uc包主程序)
│
├──Makefile (ucmain的makefile)（Makefile 可以简单的认为是一个工程文件的编译规则，描述了整个工程的编译和链接等规则。）
│
├──ucmain
│
├──src
│   │
│   └─uc	 (包含uc包的go源码)
│      │
│      ├─uc.go
│      │
│      ├─uc_test.go
│      │
│      ├─Makefile (包的makefile)
│      │
│      ├─uc.a
│      │
│      ├─_obj
│      │  │
│      │  └──uc.a
│      │
│      └─_test
│         │
│         └───uc.a
│
├──bin		(包含最终的执行文件)
│   │
│   └──ucmain
│
└──pkg
    │
    └───linux_amd64
         │
         └────uc.a	(包的目标文件)
```

src ：项目的源代码

pkg ：编译后的生成文件

bin ： 编译后的可执行文件

你的项目将作为 src 的子目录。uc 包中的功能在 uc.go 中实现。

包通常附带一个或多个测试文件，在这我们创建了一个 uc_test.go 文件

示例 9.7 test.go

```go
package uc
import "testing"

type ucTest struct {
	in, out string
}

var ucTests = []ucTest {
	ucTest{"abc", "ABC"},
	ucTest{"cvo-az", "CVO-AZ"},
	ucTest{"Antwerp", "ANTWERP"},
}

func TestUC(t *testing.T) {
	for _, ut := range ucTests {
		uc := UpperCase(ut.in)
		if uc != ut.out {
			t.Errorf("UpperCase(%s) = %s, must be %s", ut.in, uc,
			ut.out)
		}
	}
}
```

pkg ：编译后的生成文件

通过指令编译并安装包到本地：`go install uc`, 这会将 uc.a 复制到 pkg/linux_amd64 下面。
