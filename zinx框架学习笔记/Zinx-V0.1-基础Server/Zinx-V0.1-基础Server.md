# 创建zinx框架，ziface和znet模块

src ：项目的源代码

pkg ：编译后的生成文件

bin ： 编译后的可执行文件

## zinx框架与ziface和znet模块的介绍

`zinx`框架 由`ziface`模块和`znet`模块 构成。

1. `ziface`模块 存放`zinx`框架全部的抽象层接口类（struct），`zinx`中最基本的服务类接口`iserver.go`就定义在其中。
2. `znet`模块 存放网络相关的模块，用来实现网络相关的功能。

# zinx框架、ziface、znet模块位置安排与原理

1. **创建zinx框架：**

   在$GOPATH/src下创建`zinx`文件夹

   

   <u>**事实上并不必要，因为有go mod直接命令一整个文件夹想放哪就放哪**</u>

   

   src ：项目的源代码

   （pkg ：编译后的生成文件）

   （bin ： 编译后的可执行文件）

2. ##### **创建ziface、znet模块**

   zinx文件夹下创建ziface文件夹和znet文件夹。

3. **在ziface文件夹下创建模块抽象层iserver.go**

   也就是上面提到的——”`ziface`模块 存放`zinx`框架全部的抽象层接口类（struct），`zinx`中最基本的服务类接口`iserver`就定义在其中“。

   
   
   下面的代码是一个最基本的server模块该包括的要素:
   
   方法：启动/运行业务/停止 服务器的方法。
   
   属性：名称、监听的ip、监听的端口。
   
   ```go
   package ziface //定义ziface这个包。
   
   	//定义服务器接口
   type IServer interface{
       //启动服务器方法
       Start()
       //停止服务器方法
       Stop()
       //开启业务服务方法
       Serve()
   }
   ```
   
   此时文件路径：
   
   ```
   └── zinx
       ├── ziface //专门存放接口，代表抽象层
       │   └── iserver.go
       └── znet   //代表实体层
           ├──server.go    
   ```
   
4. **在zent下实现服务模块sever.go**

   `znet`模块 存放网络相关的模块，用来实现网络相关的功能。

   **下面会做出来的 <u>客户端</u> 也会放在这里。**
   
   ```go
   package znet
   
   import (
       "fmt"
       "net"
       "time"
       "zinx/ziface"
   )
   
   //iServer 接口实现，定义一个Server服务类
   type Server struct {
       //服务器的名称
       Name string
       //tcp4 or other
       IPVersion string
       //服务绑定的IP地址
       IP string
       //服务绑定的端口
       Port int
   }
   
   
   //============== 实现 ziface.IServer 里的全部接口方法 ========
   
   //开启网络服务
   func (s *Server) Start() {
       fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n", s.IP, s.Port)
   
       //开启一个go去做服务端Linster业务
       go func() {
           //1 获取一个TCP的Addr
           addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
           if err != nil {
               fmt.Println("resolve tcp addr err: ", err)
               return
           }
   
           //2 监听服务器地址
           listenner, err:= net.ListenTCP(s.IPVersion, addr)
           if err != nil {
               fmt.Println("listen", s.IPVersion, "err", err)
               return
           }
   
           //已经监听成功
           fmt.Println("start Zinx server  ", s.Name, " succ, now listenning...")
   
           //3 启动server网络连接业务
           for {
               //3.1 阻塞等待客户端建立连接请求
               conn, err := listenner.AcceptTCP()
               if err != nil {
                   fmt.Println("Accept err ", err)
                   continue
               }
   
               //3.2 TODO Server.Start() 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
   
               //3.3 TODO Server.Start() 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
   
               //我们这里暂时做一个最大512字节的回显服务
               go func () {
                   //不断的循环从客户端获取数据
                   for  {
                       buf := make([]byte, 512)
                       cnt, err := conn.Read(buf)
                       if err != nil {
                           fmt.Println("recv buf err ", err)
                           continue
                       }
                       //回显
                       if _, err := conn.Write(buf[:cnt]); err !=nil {
                           fmt.Println("write back buf err ", err)
                           continue
                       }
                   }
               }()
           }
       }()
   }
   
   func (s *Server) Stop() {
       fmt.Println("[STOP] Zinx server , name " , s.Name)
   
       //TODO  Server.Stop() 将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
   }
   
   func (s *Server) Serve() {
       s.Start()
   
       //TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加
   
   
       //阻塞,否则主Go退出， listenner的go将会退出
       for {
           time.Sleep(10*time.Second)
       }
   }
   
   
   /*
     创建一个服务器句柄
    */
   func NewServer (name string) ziface.IServer {
       s:= &Server {
           Name :name,
           IPVersion:"tcp4",
           IP:"0.0.0.0",
           Port:7777,
       }
   
       return s
   }
   ```

   此时文件路径：
   
   ```
   └── zinx
       ├── ziface         //存放`zinx`框架全部的抽象层接口类（struct）
       │   └── iserver.go //`zinx`中最基本的服务类接口`iserver`
       └── znet      //`znet`模块 存放网络相关的模块，用来实现网络相关的功能。
           ├──server.go   //一个在znet下实现的名为server.go的服务模块
   ```

   

   
   
   ## sever.go的大致思路：**<u>以后复习的时候总结</u>**
   
   1. 定义关于server的struct
   2. **<u>以后复习的时候总结</u>**



# 另外的测试方法：

当然，如果感觉go test 好麻烦，那么我们可以完全基于zinx写两个应用程序，Server.go , Client.go

Server.go

```go
package main

import (
    "zinx/znet" //用的是zinx架构里的znet模块，也就是存放网络相关的模块，用来实现网络相关的功能。
)

//Server 模块的测试函数
func main() {

    //1 创建一个server 句柄 s
    s := znet.NewServer("[zinx V0.1]")
    //为Server中的name赋值为[zinx V0.1]
/*
上面对标的是server.go中的初始化server模块的方法 :
func NewServer(name(这里就是输入值) string) ziface.IServer {
	s := &Server{
		Name:      name（这里就是输入值）,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
	}

	return s
}
*/
    //2 开启服务
    s.Serve() 
    /*
    此处s可选s.Serve(),s.Start(),s.Stop
    这都是出现在ziface里接口IServer里的方法。
    */
}
//总结一下就是先为结构体（其中蕴含着服务器的基本参数）赋予一个名字（创造句柄），再调用ziface里的接口中的方法。
```

启动Server.go

```bash
go run Server.go
```

Client.go

```go
package main

import (
    "fmt"
    "net"
    "time"
)

func main() {

    fmt.Println("Client Test ... start")
    //3秒之后发起测试请求，给服务端开启服务的机会
    time.Sleep(3 * time.Second)

    conn,err := net.Dial("tcp", "127.0.0.1:7777")
    /*调用方式
    1) TCP 连接
conn, err := net.Dial("tcp", "192.168.10.10:80")

	2) UDP 连接
conn, err := net.Dial("udp", "192.168.10.10:8888")
    */
    if err != nil {
        fmt.Println("client start err, exit!")
        return
    }
//开始不断请求
    for {
        _, err := conn.Write([]byte("hahaha")) //连接成功后对于切片的写入。功能。
        if err !=nil {
            fmt.Println("write error err ", err)
            return
        }

        buf :=make([]byte, 512) //buf是缓冲区的意思。此处是512字节的切片。
        cnt, err := conn.Read(buf) //连接成功后对于缓冲区切片的读取。
        
        if err != nil {
            fmt.Println("read buf error ")
            return
        }

        fmt.Printf(" server call back : %s, cnt = %d\n", buf,  cnt)

        time.Sleep(1*time.Second)
    }
}
```

启动Client.go进行测试

```bash
go run Client.go
```

