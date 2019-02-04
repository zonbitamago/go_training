# the_empty_interface

[https://go-tour-jp.appspot.com/methods/14](https://go-tour-jp.appspot.com/methods/14)

何もメソッドを指定していないインターフェースを、空のインターフェースと言う。

```go
interface{}
```

空のインターフェースは任意の型を保持できる。

空のインターフェースは未知の型を扱うコードで使用される。例えば、`fmt.Print`は`interface{}`型の任意の数の引数を受け取る。
