# HTTP协议详解
HTTP协议是无状态的，
- 无状态是指协议对于事务处理没有记忆能力，服务器不知道客户端是什么状态。
- 从另一方面讲，打开一个服务器上的网页和你之前打开这个服务器上的网页之间没有任何联系。
为了解决这个问题，Web程序引入了Cookie机制来维护连接的可持续状态。


## HTTP请求包（浏览器信息）
Request包分为3部分：
1. 第一部分叫Request line（请求行）, 
2. 第二部分叫Request header（请求头）,
3. *header和body之间有个空行*
4. 第三部分是body（主体）。

```
GET /domains/example/ HTTP/1.1		//请求行: 请求方法(get) 请求URI HTTP协议/协议版本
Host：www.iana.org				    //服务端的主机名
User-Agent：Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.4 (KHTML, like Gecko) Chrome/22.0.1229.94 Safari/537.4			//浏览器信息
Accept：text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8	//客户端能接收的MIME
Accept-Encoding：gzip,deflate,sdch		//是否支持流压缩
Accept-Charset：UTF-8,*;q=0.5		//客户端字符编码集
//空行,用于分割请求头和消息体
//消息体body,请求资源参数,例如POST传递的参数
```

一个URL地址用于描述一个网络上的资源。
HTTP中的GET, POST, PUT, DELETE就对应着对这个资源的查，增，改，删4个操作。
我们最常见的就是GET和POST了。
GET一般用于获取/查询资源信息，而POST一般用于更新资源信息。

GET和POST的区别:

1.  我们可以看到GET请求消息体为空，POST请求带有消息体。
2.  GET提交的数据会放在URL之后，以`?`分割URL和传输数据，参数之间以`&`相连，如`EditPosts.aspx?name=test1&id=123456`。POST方法是把提交的数据放在HTTP包的body中。
3.  GET提交的数据大小有限制（因为浏览器对URL的长度有限制），而POST方法提交的数据没有限制。
4.  GET方式提交数据，会带来安全问题，比如一个登录页面，通过GET方式提交数据时，用户名和密码将出现在URL上，如果页面可以被缓存或者其他人可以访问这台机器，就可以从历史记录获得该用户的账号和密码。

## HTTP响应包（服务器信息）
结构如下：
```
HTTP/1.1 200 OK						//状态行
Server: nginx/1.0.8					//服务器使用的WEB软件名及版本
Date:Date: Tue, 30 Oct 2012 04:14:25 GMT		//发送时间
Content-Type: text/html				//服务器发送信息的类型
Transfer-Encoding: chunked			//表示发送HTTP包是分段发的
Connection: keep-alive				//保持连接状态
Content-Length: 90					//主体内容长度
//空行 用来分割消息头和主体
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"... //消息体
```
Response包中的第一行叫做状态行，由HTTP协议版本号， 状态码， 状态消息 三部分组成。
**状态码**用来告诉HTTP客户端,HTTP服务器是否产生了预期的Response。HTTP/1.1协议中定义了5类状态码， 状态码由三位数字组成，第一个数字定义了响应的类别

-   1XX 提示信息 - 表示请求已被成功接收，继续处理
-   2XX 成功 - 表示请求已被成功接收，理解，接受
-   3XX 重定向 - 要完成请求必须进行更进一步的处理
-   4XX 客户端错误 - 请求有语法错误或请求无法实现
-   5XX 服务器端错误 - 服务器未能实现合法的请求

HTTP是一个无状态的面向连接的协议，无状态不代表HTTP不能保持TCP连接，更不能代表HTTP使用的是UDP协议（面对无连接）。

从HTTP/1.1起，默认都开启了Keep-Alive保持连接特性，简单地说，当一个网页打开完成后，客户端和服务器之间用于传输HTTP数据的TCP连接不会关闭，如果客户端再次访问这个服务器上的网页，会继续使用这一条已经建立的TCP连接。

Keep-Alive不会永久保持连接，它有一个保持时间，可以在不同服务器软件（如Apache）中设置这个时间。

就像是学react的时候，
第一次请求url，服务器端返回的是html页面，然后浏览器开始渲染HTML：当解析到HTML DOM里面的图片连接，css脚本和js脚本的链接，浏览器就会自动发起一个请求静态资源的HTTP请求，获取相对应的静态资源，然后浏览器就会渲染出来，最终将所有资源整合、渲染，完整展现在我们面前的屏幕上。