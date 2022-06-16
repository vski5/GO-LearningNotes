MySQL用 `go-sql-driver/mysql`包操作，
postgresql用`SQLC`包操作。
```sh
go get -u github.com/go-sql-driver/mysql

```

```go
import "database/sql"
import _ "go-sql-driver/mysql"
```

配置数据库连接(始终检查错误)
```go
db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")
```
