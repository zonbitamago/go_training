# bufferd_channels

[https://go-tour-jp.appspot.com/concurrency/3](https://go-tour-jp.appspot.com/concurrency/3)

チャネルは、バッファ(buffer)として使うことが出来る。バッファを持つチャネルを初期化するには、`make`の2つめの引数にバッファの長さを与える。

```go
ch := make(chan int, 100)
```

バッファが詰まったときは、チャネルへの送信をブロックする。バッファが空のときは、チャネルの受信をブロックする。
