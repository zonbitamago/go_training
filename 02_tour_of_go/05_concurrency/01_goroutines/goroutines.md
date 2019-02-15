# goroutines

[https://go-tour-jp.appspot.com/concurrency/1](https://go-tour-jp.appspot.com/concurrency/1)

**Goroutine**(ゴルーチン)は、Goのランタイムに管理される軽量なスレッドである。

```go
 go f(x, y, z)
```

と書けば、新しいgoroutineが実行される。

```go
f(x, y, z)
```

`f, x, y, z`の評価は、実行元(current)のgoroutineで管理され、`f`の実行は、新しいgoroutineで実行される。

goroutineは、同じアドレス空間で実行されるため、共有メモリへのアクセスは必ず同期する必要がある。`sync`パッケージは同期する際に役立つ方法を提供するが、別の方法があるためそれほど必要はない。