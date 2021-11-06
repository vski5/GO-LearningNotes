# go的打包成exe文件

1. 将go项目命名，这里假设为name

2. 在name这个文件所在的文件夹，右键唤出来cmd中打开，cmd里面输入`go mod init name`，项目名字和`go mod init`后面一格里的名字要一样。

   这里的cmd是在这个文件夹里操作，还没有摸索出来怎么用cd加上文件夹位置的方式，也有可能是我没有理解cd的意思。

3. 在cmd里面输入`go build name`,就能打包成exe。在hello所在的目录下用步骤2的cmd执行：`go build`

   或者在其他目录执行以下命令：`go build name`

   go编译器会去目录下查找你要编译的名为name的项目，

   大概是第二步创建的go.mod让系统知道了这个名为name的文件的位置存在。

