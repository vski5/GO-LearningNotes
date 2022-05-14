# GOLANG构建安全服务器。
基于TLS/SSL。
`NaCl`是Go的加密工具集，不读源码了，原理就是公私钥那一套，直接用就完事了。

[在 Golang 中创建安全服务器 (austburn.me)](https://austburn.me/blog/golang-server.html)

安全连接的一般流程：
-   连接到服务器。
-   在服务器和客户端上生成密钥对。
-   交换公钥。
-   开始客户端和服务器之间的通信。

在大多数情况下，服务器和客户端的行为是相同的。
不同之处在于，在服务器和客户端上开始读/写循环之前，我们执行握手以交换公钥。

先写个基础的，然后再迭代到加密，对着[[go-90-GO高级编程-1500-网络、模板与网页应用]]里例子抄就行。

直接⏩快进到用于加密的部分：

在服务器和客户端上开始读/写循环之前，执行握手以交换公钥。

在服务器和客户端上启动连接的方式是相同的。

函数生成一个密钥对，将其公钥写入连接，并从缓冲区中读取对等体的密钥。


服务器：
```go
func handleConnection(conn *net.TCPConn) { 
	sharedKey := Handshake(conn) 
	secureConnection := SecureConnection{conn: conn, sharedKey: sharedKey} // Read/write loop 
	}
```
客户端：
```go
func (c *Client) Connect() error { 
// Connect to server... 
	sharedKey := Handshake(conn) 
	secureConnection := SecureConnection{conn: conn, sharedKey: sharedKey} 
// Read/write loop 
}
```
加密：
```go
func Handshake(conn *net.TCPConn) *[32] byte {
    var peerKey, sharedKey [32]byte

    publicKey, privateKey, _ := box.GenerateKey(rand.Reader)

    // Deliver the public key
    conn.Write(publicKey[:])

    // Receive the peer key
    peerKeyArray := make([]byte, 32)
    conn.Read(peerKeyArray)

    copy(peerKey[:], peerKeyArray)

    box.Precompute(&sharedKey, &peerKey, privateKey)

    return &sharedKey
}
```
