# default_selection

[https://go-tour-jp.appspot.com/concurrency/6](https://go-tour-jp.appspot.com/concurrency/6)

どの`case`も準備ができていないならば、`select`の中の`default`が実行される。

ブロックせずに受信するなら、`default`の`case`を利用すること。

```go
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
```
