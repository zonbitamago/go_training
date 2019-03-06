# database

## 参考

- [https://qiita.com/K_ichi/items/e8826c300e797b90e40f](https://qiita.com/K_ichi/items/e8826c300e797b90e40f)
- [https://qiita.com/tenntenn/items/dddb13c15643454a7c3b](https://qiita.com/tenntenn/items/dddb13c15643454a7c3b)
- [https://www.write-ahead-log.net/entry/2017/03/31/140335](https://www.write-ahead-log.net/entry/2017/03/31/140335)
- [http://precure-3dprinter.hatenablog.jp/entry/2018/11/22/Golang%E3%81%A7%E3%83%88%E3%83%A9%E3%83%B3%E3%82%B6%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3%E3%82%92%E4%BD%BF%E3%81%86%E8%A9%B1](http://precure-3dprinter.hatenablog.jp/entry/2018/11/22/Golang%E3%81%A7%E3%83%88%E3%83%A9%E3%83%B3%E3%82%B6%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3%E3%82%92%E4%BD%BF%E3%81%86%E8%A9%B1)

## MySQL

サンプルとして、mysqlをDockerで利用する形とする。

(起動コマンド)

```sh
docker-compose up -d
```

(停止コマンド)

```sh
docker-compose down
```

## 標準パッケージ

### MySQLドライバインストール

```sh
go get -u github.com/go-sql-driver/mysql
```

### コネクション確立

`standard/Select.go`

```go
// 第2引数の形式は "user:password@tcp(host:port)/dbname"
db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/app")
if err != nil {
    panic(err.Error())
}
defer db.Close()
```

### Select文実行&1行ずつ読み込み

`standard/Select.go`

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

### トランザクション確立

トランザクションに関しては、`db.Begin()`で開始となる。トランザクション開始時の戻り値`tx`は、`db`変数と同等の機能を保持しているので、それを利用しDB操作(CRUD)を行う。

また、トランザクション開始時に、`recover()`パターンを利用して、ロールバックの設定を行う。

`standard/Insert.go`

```go
// トランザクション開始
tx, err := db.Begin()
if err != nil {
    panic(err.Error())
}

// recoverでロールバックするように設定しておく
// recoverについては以下。
// https://blog.amedama.jp/entry/2015/10/11/123535
defer func() {
    if err := recover(); err != nil {
        fmt.Println(err)
        if err := tx.Rollback(); err != nil {
            panic(err.Error())
        }
        fmt.Println("Rollbacked")
    }
}()
```

トランザクションの終了は、必要な箇所で`tx.Commit()`を行う。

```go
// コミット
if err = tx.Commit(); err != nil {
    panic(err.Error())
}
```

### Insert文実行

Insertは前述のトランザクション`tx`を利用して実行する。

`standard/Insert.go`

```go
_, err = tx.Exec(`
    INSERT INTO users(name) VALUES(?)
`, "dummy")
if err != nil {
    panic(err.Error())
}
```