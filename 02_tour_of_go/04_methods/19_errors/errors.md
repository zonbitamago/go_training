# errors

[https://go-tour-jp.appspot.com/methods/19](https://go-tour-jp.appspot.com/methods/19)

Goのプログラムは、エラーの状態を`error`値で表現する。

`error`型は`fmt.Stringer`に似た組み込みのインターフェースである。

```go
type error interface {
    Error() string
}
```

(`fmt.Stringer`と同様に、`fmt`パッケージは、変数を文字列で出力する際に`error`インターフェースを確認する。)

関数は`error`変数を返却することがある。そして、呼び出し元はエラーが`nil`かどうかを確認することでエラーを扱う。

```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

nilの`error`は成功したことを示し、nilではない`error`は失敗したことを示す。