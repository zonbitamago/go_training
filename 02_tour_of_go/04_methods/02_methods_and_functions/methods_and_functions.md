# methods_and_functions

[https://go-tour-jp.appspot.com/methods/2](https://go-tour-jp.appspot.com/methods/2)

メソッドは、レシーバ引数を通常の引数として記述すると、関数でも記述可能である。

```go
// メソッドの場合
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ファンクションの場合
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```
