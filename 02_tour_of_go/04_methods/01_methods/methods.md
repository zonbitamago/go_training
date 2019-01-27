# methods

[https://go-tour-jp.appspot.com/methods/1](https://go-tour-jp.appspot.com/methods/1)

Goにクラスは存在しないが、型にメソッド(`method`)を定義できる。

メソッドは、特別にレシーバ(`receiver`)を関数に定義する。

レシーバは`func`キーワードとメソッド名の間に自身の引数と型を表現する。

```go
type Vertex struct {
	X, Y float64
}

// この場合「(v Vertex)」がレシーバ
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```