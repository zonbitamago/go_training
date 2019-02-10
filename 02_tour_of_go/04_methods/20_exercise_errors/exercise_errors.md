# exercise_errors

[https://go-tour-jp.appspot.com/methods/20](https://go-tour-jp.appspot.com/methods/20)

`Sqrt`関数を[以前の演習](https://go-tour-jp.appspot.com/flowcontrol/8)からコピーし、`error`の値を返すように修正すること。

'Sqrt'関数は複素数をサポートしていないので、負の数が与えられたときに、nil以外のエラー値を返す必要がある。

新しい型

```go
type ErrNegativeSqrt float64
```

を作成すること。

そして、`ErrNegativeSqrt(-2).Error()`で、`"cannot Sqrt negative number: -2"`を返すような

```go
func (e ErrNegativeSqrt) Error() string
```

メソッドを実装し、`error`インターフェースを満たすようにすること。