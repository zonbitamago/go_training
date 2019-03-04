# database

サンプルとして、mysqlをDockerで利用する形とする。

参考：[https://qiita.com/K_ichi/items/e8826c300e797b90e40f](https://qiita.com/K_ichi/items/e8826c300e797b90e40f)

(起動コマンド)

```sh
docker-compose up -d
```

(停止コマンド)

```sh
docker-compose down
```

## 標準パッケージ

参考

- [https://qiita.com/tenntenn/items/dddb13c15643454a7c3b](https://qiita.com/tenntenn/items/dddb13c15643454a7c3b)
- [https://www.write-ahead-log.net/entry/2017/03/31/140335](https://www.write-ahead-log.net/entry/2017/03/31/140335)

### MySQLドライバインストール

```sh
go get -u github.com/go-sql-driver/mysql
```

### コネクション確立

`standard/database.go`

```go
// 第2引数の形式は "user:password@tcp(host:port)/dbname"
db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/app")
if err != nil {
    panic(err.Error())
}
defer db.Close()
```

### SQL実行&1行ずつ読み込み

`standard/database.go`

```go
rows, err := db.Query(`
    SELECT
        id
        ,name
    FROM users
    WHERE id = ?
`, 1)
if err != nil {
    panic(err.Error())
}
defer rows.Close()

for rows.Next() {
    var user User
    err := rows.Scan(&(user.ID), &(user.Name))
    if err != nil {
        panic(err.Error())
    }
    fmt.Println(user)
}

if err := rows.Err(); err != nil {
    panic(err.Error())
}
```