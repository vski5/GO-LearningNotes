教程来自：[在Go |从头开始构建BitTorrent客户端李杰西 (jse.li)](https://blog.jse.li/posts/torrent/)

[用GO从零建立BitTorrent客户端 – HaoranDeng's blog – 会写点代码，会写点小说 (mynameisdhr.com)](https://blog.mynameisdhr.com/YongGOCongLingJianLiBitTorrentKeHuDuan/)


.torrent 文件描述了可下载文件的内容和连接到特定 tracker 的信息。
文件采用Bencode编码。
Bencode 编码与 JSON 的结构大致相同，先解码，读出来内容，再做下一步打算。

解码器用现成的：
[jackpal/bencode-go：一种 Go 语言绑定，用于以 BitTorrent 对等文件共享协议使用的 bencode 格式对数据进行编码和解码。 (github.com)](https://github.com/jackpal/bencode-go)

直接导入包就行。

在Bencode 编码中，我们能看到 tracker 的 URL、种子建立的时间（使用 Unix 时间戳表示）、文件的名称和大小和一个巨大的含有文件的各个 **piece** 的 SHA-1 哈希值的 BLOB（binary large object）。这个 BLOB 的大小与我们希望下载的文件的一部分是相等的。一个 piece 的实际大小根据不同的 torrents 会有一些变化，通常介于 256KB 和 1MB 之间。这意味着一个大文件可能由_数千个_ pieces 组成。我们将要从其他 peers 那里下载这些 pieces，将他们与 .torrent 文件中的哈希值进行核对，把它们组装到一起，最后我们就得到了一个完整的文件！

这一机制允许我们在下载过程中验证每一个 piece 的完整性。这让 BitTorrent 具有了抵抗文件意外损坏和大规模 **torrent poisoning** 的能力。我们一定能得到我们希望下载的文件，除非某个攻击者能够使用原像攻击（哈希不可逆以及抗碰撞性——译者注）破解 SHA-1。

磁力链接可以包括一个或多个参数，之间用'&'隔开。参数的顺序在文件在标准中没有记录。有一些参数的值对于客户端正确解析磁力链接很重要。

 >  magnet:? xl = [字节大小]& dn = [文件名（已编码URL）]& xt = urn: tree: tiger: [ TTH hash（Base32）]
 
 标准还建议同类的多个参数可以在参数名称后面加上".1", ".2"等来使用，例如

Magnet-icon.gif magnet:?xt.1=urn:sha1:YNCKHTQCWBTRNJIV4WNAE52SJUQCZO5C&xt.2=urn:sha1:TXGCZQTH26NL6OUQAJJPFALHG2LTGBC7

