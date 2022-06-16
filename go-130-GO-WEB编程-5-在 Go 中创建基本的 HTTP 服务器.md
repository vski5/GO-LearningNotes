# go-130-GO-WEB编程-5-在 Go 中创建基本的 HTTP 服务器

在Go中创建一个基本的HTTP服务器应该具备哪些功能：
- 处理动态请求。处理来自用户浏览网站、登录账户或发布图片的传入请求。
- 服务静态资产。向浏览器提供JavaScript、CSS和图像，为用户创造动态体验。
- 接受连接。HTTP服务器必须监听一个特定的端口，以便能够接受来自互联网的连接。


## 处理动态请求
"net/http"包含**接受请求**和**处理动态请求**所需的所有实用程序。
我们可以用`http.HandleFunc`这个函数注册一个新的处理程序。
- 第一个参数(`pattern string`)是需要匹配的路径，
- 第二个参数`handler func(ResponseWriter, *Request)`是需要执行的函数。
- 在这个例子中。当有人浏览你的网站时，会收到一条信息`fmt.Fprint(w, "Welcome to my website!")`。
```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to my website!")
})
```

对于动态请求，`http.Request`包含所有关于请求和它的参数的信息。
一个Request代表一个 由服务器接收的 或 将由客户端发送的 HTTP请求。

你可以用GET参数或POST参数`r.URL.Query().Get("token")` （来自HTML表单的字段）来读取 `r.FormValue("email")`

## 服务静态资产
为了提供静态资产，如JavaScript、CSS和图片，我们使用内置的`http.FileServer`将其指向一个URL路径。
为了使文件服务器正常工作，文件服务器需要知道从哪里提供文件。我们可以这样做
```go
fs := http.FileServer(http.Dir("static/"))
```
一旦我们的文件服务器到位，我们只需要将一个url路径指向它，就像我们对动态请求所做的那样。

需要注意的一点是：为了正确地提供文件，我们需要剥离url路径的一部分。通常这是我们的文件所在的目录的名称。
```go
http.Handle("/static/", http.StripPrefix("/static/", fs))
```
`func StripPrefix(prefix string, h Handler) Handler`
Strip Prefix中文意思是剥夺前缀。
函数StripPrefix返回一个处理程序，通过从请求的URL的Path（如果设置了RawPath）中 **移除给定的前缀** 并调用 处理程序`h Handler`(此处为上面设置的函数fs) 来提供HTTP请求。
前缀必须完全匹配：如果请求中的前缀包含转义字符，那么回复也是一个HTTP 404 not found错误。

一个例子：用户访问126.0.0.1/static/dir/file的时候，/static/被剥夺，也就是说，实际上访问的是服务器上的dir/file资源，

总结一下：
- 声明为文件服务器介绍文件路径的函数fs。
- 为 请求的URL的Path 注册了一个处理函数用于剥夺前缀`/static/`并且调用上面的函数fs介绍文件路径。

```go
fs := http.FileServer(http.Dir("static/"))
http.Handle("/static/", http.StripPrefix("/static/", fs))
```

## 接受连接
```go
http.ListenAndServe(":80", nil)
```

