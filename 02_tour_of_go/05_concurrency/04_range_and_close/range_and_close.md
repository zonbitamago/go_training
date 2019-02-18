# range_and_close

[https://go-tour-jp.appspot.com/concurrency/4](https://go-tour-jp.appspot.com/concurrency/4)

送り手は、これ以上送信する値が無いことを示すため、チャネルを`close`できる。受け手は、受信の式に2つ目のパラメータを割り当て、そのチャネルが`close`されているかどうかを確認できる。

```go
v, ok := <-ch
```

受信する値がない、かつ、チャネルが閉じているなら、`ok`の変数は、`false`になる。

ループの`for i := range c`は、チャネルが閉じられるまで、チャネルから繰り返し値を受信する。