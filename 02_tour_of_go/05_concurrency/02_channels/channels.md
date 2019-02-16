# channels

[https://go-tour-jp.appspot.com/concurrency/2](https://go-tour-jp.appspot.com/concurrency/2)

チャネル(*Channel*)型は、チャネルオペレータの`<-`を用いて値の送受信が出来る通り道である。

```go
ch <- v    // v をチャネル ch へ送信する
v := <-ch  // ch から受信した変数を v へ割り当てる
```

マップとスライスのように、チャネルは使う前に以下のように作成する。

```go
ch := make(chan int)
```

通常、片方が準備できるまでは送受信はブロックされる。これにより、明確なロックや条件変数がなくても、goroutineの同期を可能とする。

サンプルコードでは、スライス内の数値を合算し、2つのgoroutine間で作業を分配する。両方のgoroutineで計算が完了すると、最終結果が計算される。