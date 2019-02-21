# sync_mutex

[https://go-tour-jp.appspot.com/concurrency/9](https://go-tour-jp.appspot.com/concurrency/9)

チャネルがgoroutine間で通信する方法であることは、今までの部分でよくわかったと思う。

しかし、通信する必要がない場合や、コンフリクトを避けるために一度に一つのgoroutineだけが変数にアクセス出来るようにしたい場合はどうか？

上記のコンセプトは排他制御(mutex exclusion)と呼ばれ、このデータ構造を指す一般的な名称はmutex(ミューテックス)である。

Goの標準ライブラリは、排他制御を`sync.Mutex`と次の２つのメソッドで提供している。

```go
Lock
Unlock
```

`Inc`メソッドにあるように、`Lock`と`Unlock`で囲むことで排他制御で実行するコードを定義することが出来る。

```go
// Inc 排他制御カウンタ増加処理
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}
```

`Value`メソッドのように、mutexがUnlockされることを保証するために、`defer`を利用することも出来る。

```go
// Value 排他制御カウンタ値取得処理
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}
```
